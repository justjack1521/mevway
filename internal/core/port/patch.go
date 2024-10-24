package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
)

type PatchRepository interface {
	GetLatestPatch(ctx context.Context, environment uuid.UUID) (patch.Patch, error)
	GetPatchList(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error)
	GetOpenIssuesList(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error)
}

type PatchService interface {
	GetCurrentPatch(ctx context.Context, environment uuid.UUID) (patch.Patch, error)
	ListPatches(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error)
	ListOpenIssues(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error)
}
