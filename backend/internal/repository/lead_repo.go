package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yourusername/binghatti-backend/internal/models"
)

type LeadRepository struct {
	db     *pgxpool.Pool
	mu     sync.RWMutex
	memory []models.Lead
}

func NewLeadRepository(db *pgxpool.Pool) *LeadRepository {
	return &LeadRepository{
		db:     db,
		memory: []models.Lead{},
	}
}

func (r *LeadRepository) CreateLead(ctx context.Context, req *models.CreateLeadRequest) (*models.Lead, error) {
	leadSource := req.LeadSource
	if leadSource == "" {
		leadSource = "Website VIP Form"
	}

	lead := models.Lead{
		ID:                fmt.Sprintf("lead-%d", time.Now().UnixNano()),
		ProjectID:         req.ProjectID,
		FullName:          req.FullName,
		WhatsAppNumber:    req.WhatsAppNumber,
		Email:             req.Email,
		BudgetRange:       req.BudgetRange,
		InvestmentPurpose: req.InvestmentPurpose,
		LeadSource:        leadSource,
		Status:            "new",
		CreatedAt:         time.Now(),
	}

	if r.db != nil {
		query := `
			INSERT INTO leads (project_id, full_name, whatsapp_number, email, budget_range, investment_purpose, lead_source)
			VALUES ($1, $2, $3, $4, $5, $6, COALESCE(NULLIF($7, ''), 'Website VIP Form'))
			RETURNING id, project_id, full_name, whatsapp_number, email, budget_range, investment_purpose, lead_source, status, created_at;
		`
		err := r.db.QueryRow(ctx, query,
			req.ProjectID,
			req.FullName,
			req.WhatsAppNumber,
			req.Email,
			req.BudgetRange,
			req.InvestmentPurpose,
			req.LeadSource,
		).Scan(
			&lead.ID,
			&lead.ProjectID,
			&lead.FullName,
			&lead.WhatsAppNumber,
			&lead.Email,
			&lead.BudgetRange,
			&lead.InvestmentPurpose,
			&lead.LeadSource,
			&lead.Status,
			&lead.CreatedAt,
		)
		if err != nil {
			fmt.Printf("Warning DB insert lead: %v. Saving in memory.\n", err)
		}
	}

	r.mu.Lock()
	r.memory = append([]models.Lead{lead}, r.memory...)
	r.mu.Unlock()

	return &lead, nil
}

func (r *LeadRepository) GetAllLeads(ctx context.Context) ([]models.Lead, error) {
	if r.db != nil {
		query := `
			SELECT id, project_id, full_name, whatsapp_number, email, budget_range, investment_purpose, lead_source, status, created_at
			FROM leads
			ORDER BY created_at DESC;
		`
		rows, err := r.db.Query(ctx, query)
		if err == nil {
			defer rows.Close()
			var leads []models.Lead
			for rows.Next() {
				var l models.Lead
				if err := rows.Scan(
					&l.ID,
					&l.ProjectID,
					&l.FullName,
					&l.WhatsAppNumber,
					&l.Email,
					&l.BudgetRange,
					&l.InvestmentPurpose,
					&l.LeadSource,
					&l.Status,
					&l.CreatedAt,
				); err == nil {
					leads = append(leads, l)
				}
			}
			if len(leads) > 0 {
				return leads, nil
			}
		}
	}

	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.memory, nil
}

func (r *LeadRepository) UpdateLeadStatus(ctx context.Context, id string, status string) error {
	r.mu.Lock()
	for i, l := range r.memory {
		if l.ID == id {
			r.memory[i].Status = status
			break
		}
	}
	r.mu.Unlock()

	if r.db != nil {
		query := `UPDATE leads SET status = $1 WHERE id = $2;`
		_, err := r.db.Exec(ctx, query, status, id)
		if err != nil {
			fmt.Printf("Warning DB update lead status: %v\n", err)
		}
	}

	return nil
}
