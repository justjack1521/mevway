package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
)

type PatchRepository interface {
	GetLatestPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error)
	GetPatchList(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error)
	GetAllowedPatchList(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error)
	GetOpenIssuesList(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error)
	GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error)
	GetTopLevelIssueList(ctx context.Context) ([]patch.Issue, error)
}

type PatchService interface {
	GetCurrentPatch(ctx context.Context, application string, environment uuid.UUID) (patch.Patch, error)
	ListPatches(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error)
	ListAllowPatches(ctx context.Context, application string, environment uuid.UUID) ([]patch.Patch, error)
	ListOpenIssues(ctx context.Context, environment uuid.UUID) ([]patch.KnownIssue, error)
	ListTopIssues(ctx context.Context) ([]patch.Issue, error)
	GetIssue(ctx context.Context, id uuid.UUID) (patch.Issue, error)
}
