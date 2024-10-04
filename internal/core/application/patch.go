package application

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/port"
	"mevway/internal/domain/patch"
)

type PatchService struct {
	repository port.PatchRepository
}

func NewPatchService(repository port.PatchRepository) *PatchService {
	return &PatchService{repository: repository}
}

func (s *PatchService) Get(ctx context.Context, environment uuid.UUID) (patch.Patch, error) {
	return s.repository.GetLatest(ctx, environment)
}

func (s *PatchService) GetList(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error) {
	return s.repository.GetList(ctx, environment, limit)
}
