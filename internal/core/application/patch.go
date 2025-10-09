package application

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
	"mevway/internal/core/port"
)

type PatchService struct {
	repository port.PatchRepository
}

func NewPatchService(repository port.PatchRepository) *PatchService {
	return &PatchService{repository: repository}
}

func (s *PatchService) GetCurrentPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error) {
	return s.repository.GetLatestPatch(ctx, application, environment)
}

func (s *PatchService) ListPatches(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error) {
	return s.repository.GetPatchList(ctx, environment, limit)
}

func (s *PatchService) ListAllowPatches(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error) {
	return s.repository.GetAllowedPatchList(ctx, application, environment)
}

func (s *PatchService) ListOpenIssues(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error) {
	return s.repository.GetOpenIssuesList(ctx, environment)
}

func (s *PatchService) ListTopIssues(ctx context.Context) ([]patch.Issue, error) {
	return s.repository.GetTopLevelIssueList(ctx)
}
