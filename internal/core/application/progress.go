package application

import (
	"context"
	"mevway/internal/core/domain/content"
	"mevway/internal/core/port"
)

type ProgressService struct {
	repository port.ProgressRepository
}

func (s *ProgressService) ListRelease(ctx context.Context) ([]content.GameFeatureRelease, error) {
	return s.repository.GetReleaseList(ctx)
}

func (s *ProgressService) ListProgress(ctx context.Context) ([]content.GameFeature, error) {
	return s.repository.GetProgressList(ctx)
}

func NewProgressService(repository port.ProgressRepository) *ProgressService {
	return &ProgressService{repository: repository}
}
