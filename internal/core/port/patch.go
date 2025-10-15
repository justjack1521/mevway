package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
)

type PatchRepository interface {
	GetLatestPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error)
	GetPatchList(ctx context.Context, environment uuid.UUID, offset, limit int) ([]patch.Patch, error)
	GetAllowedPatchList(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error)
}

type PatchService interface {
	GetCurrentPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error)
	ListPatches(ctx context.Context, environment uuid.UUID, offset, limit int) ([]patch.Patch, error)
	ListAllowPatches(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error)
}
