package port

import (
	"context"
	"mevway/internal/core/domain/patch"

	uuid "github.com/satori/go.uuid"
)

type PatchRepository interface {
	GetLatestPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error)
	GetPatchListCount(ctx context.Context) (int, error)
	GetPatchList(ctx context.Context, environment uuid.UUID, offset, limit int) ([]patch.Patch, error)
	GetAllowedPatchList(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error)
	GetAllPatchVersionList(ctx context.Context, application string) ([]string, error)
}

type PatchService interface {
	GetCurrentPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error)
	ListPatches(ctx context.Context, environment uuid.UUID, offset, limit int) ([]patch.Patch, error)
	ListAllowPatches(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error)
	ListPatchCount(ctx context.Context) (int, error)
	ListALlPatchVersions(ctx context.Context, application string) ([]string, error)
}
