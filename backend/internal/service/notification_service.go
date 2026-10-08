package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/yourusername/binghatti-backend/internal/models"
)

type NotificationService struct {
	httpClient       *http.Client
	whatsappToken    string
	whatsappPhoneID  string
	externalWebhook  string
	agentAlertURL    string
}

func NewNotificationService() *NotificationService {
	return &NotificationService{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		whatsappToken:   os.Getenv("WHATSAPP_CLOUD_API_TOKEN"),
		whatsappPhoneID: os.Getenv("WHATSAPP_PHONE_NUMBER_ID"),
		externalWebhook: os.Getenv("EXTERNAL_CRM_WEBHOOK_URL"),
		agentAlertURL:   os.Getenv("AGENT_ALERT_WEBHOOK_URL"),
	}
}

// SendLeadNotificationsAsync dispatches background tasks without blocking the HTTP response
func (s *NotificationService) SendLeadNotificationsAsync(lead *models.Lead) {
	go func() {
		// Create detached background context to avoid HTTP request cancellation
		ctx := context.Background()

		log.Printf("🔔 Processing asynchronous notifications for Lead ID: %s (%s)", lead.ID, lead.FullName)

		// 1. Send WhatsApp confirmation to VIP Lead
		if err := s.sendWhatsAppMessageWithRetry(ctx, lead); err != nil {
			log.Printf("❌ Failed to send WhatsApp message to %s after retries: %v", lead.WhatsAppNumber, err)
		} else {
			log.Printf("✅ WhatsApp confirmation dispatched to %s", lead.WhatsAppNumber)
		}

		// 2. Send Internal Agent Alert Notification
		if err := s.sendInternalAgentAlertWithRetry(ctx, lead); err != nil {
			log.Printf("❌ Failed to send Internal Agent Alert for lead %s: %v", lead.ID, err)
		} else {
			log.Printf("✅ Internal agent notification dispatched for lead %s", lead.ID)
		}

		// 3. Send External CRM Webhook (Google Sheets / HubSpot)
		if err := s.sendExternalWebhookWithRetry(ctx, lead); err != nil {
			log.Printf("❌ Failed to send External CRM Webhook for lead %s: %v", lead.ID, err)
		} else {
			log.Printf("✅ External CRM webhook delivered for lead %s", lead.ID)
		}
	}()
}

// 1. WhatsApp Cloud API Payload & Retry Logic (Meta WhatsApp Business API format)
type whatsappTextPayload struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

func (s *NotificationService) sendWhatsAppMessageWithRetry(ctx context.Context, lead *models.Lead) error {
	if s.whatsappToken == "" || s.whatsappToken == "your_whatsapp_bearer_token_here" || s.whatsappPhoneID == "" || s.whatsappPhoneID == "your_whatsapp_phone_number_id_here" {
		log.Printf("ℹ️ WhatsApp Cloud API credentials not configured (development mode). Skipping Meta API call for %s.", lead.WhatsAppNumber)
		return nil
	}

	messageBody := fmt.Sprintf(
		"Hello %s,\n\nThank you for registering your interest in Binghatti Luxury Residences! 🏰\n\nOur VIP Advisory Team has received your inquiry for budget %s (%s).\nAn advisor will reach out to you shortly via WhatsApp with exclusive floor plans and starting prices.\n\nWarm regards,\nBinghatti Sales Team",
		lead.FullName,
		lead.BudgetRange,
		lead.InvestmentPurpose,
	)

	payload := whatsappTextPayload{
		MessagingProduct: "whatsapp",
		To:               lead.WhatsAppNumber,
		Type:             "text",
	}
	payload.Text.Body = messageBody

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal WhatsApp payload: %w", err)
	}

	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", s.whatsappPhoneID)

	return s.executeHTTPWithRetry(ctx, "WhatsApp API", "POST", url, jsonBytes, map[string]string{
		"Authorization": "Bearer " + s.whatsappToken,
		"Content-Type":  "application/json",
	})
}

// 2. Internal Agent Alert Retry Logic
func (s *NotificationService) sendInternalAgentAlertWithRetry(ctx context.Context, lead *models.Lead) error {
	url := s.agentAlertURL
	if url == "" || strings.Contains(url, "your-internal-alert-endpoint.com") || strings.Contains(url, "example.com") {
		log.Printf("ℹ️ Agent Alert Webhook URL not configured (placeholder mode). Skipping agent notification.")
		return nil
	}

	alertPayload := map[string]any{
		"event":              "NEW_VIP_LEAD_REGISTERED",
		"lead_id":            lead.ID,
		"full_name":          lead.FullName,
		"whatsapp":           lead.WhatsAppNumber,
		"email":              lead.Email,
		"budget_range":       lead.BudgetRange,
		"investment_purpose": lead.InvestmentPurpose,
		"source":             lead.LeadSource,
		"timestamp":          lead.CreatedAt.Format(time.RFC3339),
	}

	jsonBytes, err := json.Marshal(alertPayload)
	if err != nil {
		return err
	}

	return s.executeHTTPWithRetry(ctx, "Agent Alert Webhook", "POST", url, jsonBytes, map[string]string{
		"Content-Type": "application/json",
	})
}

// 3. External CRM / Google Sheets Webhook Retry Logic
func (s *NotificationService) sendExternalWebhookWithRetry(ctx context.Context, lead *models.Lead) error {
	url := s.externalWebhook
	if url == "" || strings.Contains(url, "your-crm-webhook-endpoint.com") || strings.Contains(url, "example.com") {
		log.Printf("ℹ️ External CRM Webhook URL not configured (placeholder mode). Skipping CRM webhook.")
		return nil
	}

	jsonBytes, err := json.Marshal(lead)
	if err != nil {
		return err
	}

	return s.executeHTTPWithRetry(ctx, "External CRM Webhook", "POST", url, jsonBytes, map[string]string{
		"Content-Type": "application/json",
	})
}

// Helper: Exponential Backoff HTTP Execution (Retries up to 3 times: 1s, 2s, 4s)
func (s *NotificationService) executeHTTPWithRetry(ctx context.Context, serviceName, method, url string, payload []byte, headers map[string]string) error {
	maxRetries := 3
	backoff := 1 * time.Second

	var lastErr error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
		if err != nil {
			return err
		}

		for k, v := range headers {
			req.Header.Set(k, v)
		}

		resp, err := s.httpClient.Do(req)
		if err == nil {
			defer resp.Body.Close()
			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				return nil // Success!
			}
			lastErr = fmt.Errorf("http status code %d", resp.StatusCode)
		} else {
			lastErr = err
		}

		log.Printf("⚠️ [%s] Attempt %d/%d failed: %v. Retrying in %v...", serviceName, attempt, maxRetries, lastErr, backoff)
		time.Sleep(backoff)
		backoff *= 2 // Exponential backoff (1s -> 2s -> 4s)
	}

	return fmt.Errorf("[%s] exhausted all %d retries. Last error: %w", serviceName, maxRetries, lastErr)
}
