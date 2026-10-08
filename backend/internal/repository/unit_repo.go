package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/yourusername/binghatti-backend/internal/models"
)

type UnitRepository struct {
	db *pgxpool.Pool
}

func NewUnitRepository(db *pgxpool.Pool) *UnitRepository {
	return &UnitRepository{db: db}
}

func (r *UnitRepository) GetUnitsByProjectID(ctx context.Context, projectID string) ([]models.Unit, error) {
	query := `
		SELECT id, project_id, unit_type, size_sqft, price_aed, floor_plan_image_url, floor_plan_pdf_url, status, created_at
		FROM units
		WHERE project_id = $1
		ORDER BY price_aed ASC;
	`
	rows, err := r.db.Query(ctx, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to query units: %w", err)
	}
	defer rows.Close()

	var units []models.Unit
	for rows.Next() {
		var u models.Unit
		err := rows.Scan(
			&u.ID,
			&u.ProjectID,
			&u.UnitType,
			&u.SizeSqft,
			&u.PriceAED,
			&u.FloorPlanImageURL,
			&u.FloorPlanPDFURL,
			&u.Status,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan unit row: %w", err)
		}
		units = append(units, u)
	}
	return units, nil
}
