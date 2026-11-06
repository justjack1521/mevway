package application

import (
	"context"
	"mevway/internal/core/domain/patch"
	"mevway/internal/core/port"

	uuid "github.com/satori/go.uuid"
)

const (
	maxListLimit = 10
)

type PatchService struct {
	repository port.PatchRepository
}

func (s *PatchService) ListALlPatchVersions(ctx context.Context, application string) ([]string, error) {
	return s.repository.GetAllPatchVersionList(ctx, application)
}

func NewPatchService(repository port.PatchRepository) *PatchService {
	return &PatchService{repository: repository}
}

func (s *PatchService) GetCurrentPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error) {
	return s.repository.GetLatestPatch(ctx, application, environment)
}

func (s *PatchService) ListPatches(ctx context.Context, environment uuid.UUID, offset, limit int) ([]patch.Patch, error) {

	if limit < 0 || limit > maxListLimit {
		limit = maxListLimit
	}

	if offset < 0 {
		offset = 0
	}

	return s.repository.GetPatchList(ctx, environment, offset, limit)
}

func (s *PatchService) ListPatchCount(ctx context.Context) (int, error) {
	return s.repository.GetPatchListCount(ctx)
}

func (s *PatchService) ListAllowPatches(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error) {
	return s.repository.GetAllowedPatchList(ctx, application, environment)
}

// IsPatchAllowed reports whether a client on the given patch may connect - the same list
// /patch/allow answers from.
func (s *PatchService) IsPatchAllowed(ctx context.Context, application string, id uuid.UUID) (bool, error) {
	allowed, err := s.repository.GetAllowedPatchList(ctx, application, uuid.Nil)
	if err != nil {
		return false, err
	}
	for _, value := range allowed {
		if uuid.Equal(value.SysID, id) {
			return true, nil
		}
	}
	return false, nil
}
