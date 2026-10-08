package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yourusername/binghatti-backend/internal/models"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	query := `
		SELECT id, slug, title, tagline, location_name, starting_price_aed, handover_date, description, brochure_url, hero_video_url, created_at
		FROM projects
		ORDER BY created_at DESC;
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var p models.Project
		err := rows.Scan(
			&p.ID,
			&p.Slug,
			&p.Title,
			&p.Tagline,
			&p.LocationName,
			&p.StartingPriceAED,
			&p.HandoverDate,
			&p.Description,
			&p.BrochureURL,
			&p.HeroVideoURL,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project row: %w", err)
		}
		projects = append(projects, p)
	}
	return projects, nil
}

func (r *ProjectRepository) GetProjectBySlug(ctx context.Context, slug string) (*models.Project, error) {
	query := `
		SELECT id, slug, title, tagline, location_name, starting_price_aed, handover_date, description, brochure_url, hero_video_url, created_at
		FROM projects
		WHERE slug = $1;
	`
	var p models.Project
	err := r.db.QueryRow(ctx, query, slug).Scan(
		&p.ID,
		&p.Slug,
		&p.Title,
		&p.Tagline,
		&p.LocationName,
		&p.StartingPriceAED,
		&p.HandoverDate,
		&p.Description,
		&p.BrochureURL,
		&p.HeroVideoURL,
		&p.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, fmt.Errorf("query error: %w", err)
	}
	return &p, nil
}
