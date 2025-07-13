package application

import (
	"context"
	"mevway/internal/core/domain/progress"
	"mevway/internal/core/port"
)

type ProgressService struct {
	repository port.ProgressRepository
}

func (s *ProgressService) ListProgress(ctx context.Context) ([]progress.GameFeature, error) {
	return s.repository.GetProgressList(ctx)
}

func NewProgressService(repository port.ProgressRepository) *ProgressService {
	return &ProgressService{repository: repository}
}
