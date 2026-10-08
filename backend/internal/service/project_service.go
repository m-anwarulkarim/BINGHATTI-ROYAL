package service

import (
	"context"

	"github.com/yourusername/binghatti-backend/internal/models"
	"github.com/yourusername/binghatti-backend/internal/repository"
)

type ProjectService struct {
	projectRepo *repository.ProjectRepository
	unitRepo    *repository.UnitRepository
}

func NewProjectService(pRepo *repository.ProjectRepository, uRepo *repository.UnitRepository) *ProjectService {
	return &ProjectService{
		projectRepo: pRepo,
		unitRepo:    uRepo,
	}
}

func (s *ProjectService) GetAllProjects(ctx context.Context) ([]models.Project, error) {
	return s.projectRepo.GetAllProjects(ctx)
}

func (s *ProjectService) GetProjectDetailsBySlug(ctx context.Context, slug string) (*models.ProjectWithUnits, error) {
	project, err := s.projectRepo.GetProjectBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}

	units, err := s.unitRepo.GetUnitsByProjectID(ctx, project.ID)
	if err != nil {
		units = []models.Unit{} // Empty fallback slice instead of nil
	}

	return &models.ProjectWithUnits{
		Project: *project,
		Units:   units,
	}, nil
}
