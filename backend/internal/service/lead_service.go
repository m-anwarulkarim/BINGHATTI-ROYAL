package service

import (
	"context"
	"errors"
	"strings"

	"github.com/yourusername/binghatti-backend/internal/models"
	"github.com/yourusername/binghatti-backend/internal/repository"
)

type LeadService struct {
	leadRepo     *repository.LeadRepository
	notifService *NotificationService
}

func NewLeadService(repo *repository.LeadRepository, notif *NotificationService) *LeadService {
	return &LeadService{
		leadRepo:     repo,
		notifService: notif,
	}
}

func (s *LeadService) CreateVIPLead(ctx context.Context, req *models.CreateLeadRequest) (*models.Lead, error) {
	req.FullName = strings.TrimSpace(req.FullName)
	req.WhatsAppNumber = strings.TrimSpace(req.WhatsAppNumber)
	req.Email = strings.TrimSpace(req.Email)

	if req.FullName == "" {
		return nil, errors.New("full_name is required")
	}
	if req.WhatsAppNumber == "" {
		return nil, errors.New("whatsapp_number is required")
	}
	if req.Email == "" || !strings.Contains(req.Email, "@") {
		return nil, errors.New("valid email is required")
	}

	lead, err := s.leadRepo.CreateLead(ctx, req)
	if err != nil {
		return nil, err
	}

	if s.notifService != nil {
		s.notifService.SendLeadNotificationsAsync(lead)
	}

	return lead, nil
}

func (s *LeadService) ListLeads(ctx context.Context) ([]models.Lead, error) {
	return s.leadRepo.GetAllLeads(ctx)
}

func (s *LeadService) UpdateLeadStatus(ctx context.Context, id string, status string) error {
	id = strings.TrimSpace(id)
	status = strings.TrimSpace(status)
	if id == "" {
		return errors.New("id is required")
	}
	if status == "" {
		return errors.New("status is required")
	}
	return s.leadRepo.UpdateLeadStatus(ctx, id, status)
}
