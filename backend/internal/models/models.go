package models

import "time"

// User Entity (Admin / Agent)
type User struct {
	ID           string    `json:"id"`
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

// Project Entity
type Project struct {
	ID                string    `json:"id"`
	Slug              string    `json:"slug"`
	Title             string    `json:"title"`
	Tagline           string    `json:"tagline"`
	LocationName      string    `json:"location_name"`
	StartingPriceAED  float64   `json:"starting_price_aed"`
	HandoverDate      string    `json:"handover_date"`
	Description       string    `json:"description"`
	BrochureURL       string    `json:"brochure_url"`
	HeroVideoURL      string    `json:"hero_video_url"`
	CreatedAt         time.Time `json:"created_at"`
}

// Unit Entity
type Unit struct {
	ID                string    `json:"id"`
	ProjectID         string    `json:"project_id"`
	UnitType          string    `json:"unit_type"`
	SizeSqft          float64   `json:"size_sqft"`
	PriceAED          float64   `json:"price_aed"`
	FloorPlanImageURL string    `json:"floor_plan_image_url"`
	FloorPlanPDFURL   string    `json:"floor_plan_pdf_url"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
}

// Lead Entity
type Lead struct {
	ID                string    `json:"id"`
	ProjectID         *string   `json:"project_id,omitempty"`
	FullName          string    `json:"full_name"`
	WhatsAppNumber    string    `json:"whatsapp_number"`
	Email             string    `json:"email"`
	BudgetRange       string    `json:"budget_range"`
	InvestmentPurpose string    `json:"investment_purpose"`
	LeadSource        string    `json:"lead_source"`
	Status            string    `json:"status"`
	AssignedAgentID   *string   `json:"assigned_agent_id,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
}

// ProjectWithUnits DTO for single project page
type ProjectWithUnits struct {
	Project
	Units []Unit `json:"units"`
}

// API DTOs
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type CreateLeadRequest struct {
	ProjectID         *string `json:"project_id,omitempty"`
	FullName          string  `json:"full_name"`
	WhatsAppNumber    string  `json:"whatsapp_number"`
	Email             string  `json:"email"`
	BudgetRange       string  `json:"budget_range"`
	InvestmentPurpose string  `json:"investment_purpose"`
	LeadSource        string  `json:"lead_source"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}
