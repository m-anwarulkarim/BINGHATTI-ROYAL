package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yourusername/binghatti-backend/internal/models"
	"github.com/yourusername/binghatti-backend/internal/service"
)

type LeadHandler struct {
	leadService *service.LeadService
}

func NewLeadHandler(service *service.LeadService) *LeadHandler {
	return &LeadHandler{leadService: service}
}

func (h *LeadHandler) RegisterLead(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req models.CreateLeadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "invalid JSON request payload",
		})
		return
	}

	lead, err := h.leadService.CreateVIPLead(r.Context(), &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "VIP registration submitted successfully",
		Data:    lead,
	})
}

func (h *LeadHandler) ListLeads(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	leads, err := h.leadService.ListLeads(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Data:    leads,
	})
}

type UpdateLeadStatusRequest struct {
	Status string `json:"status"`
}

func (h *LeadHandler) UpdateLeadStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "lead id is required",
		})
		return
	}

	var req UpdateLeadStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   "invalid JSON payload",
		})
		return
	}

	if err := h.leadService.UpdateLeadStatus(r.Context(), id, req.Status); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(models.APIResponse{
		Success: true,
		Message: "Lead status updated successfully",
	})
}
