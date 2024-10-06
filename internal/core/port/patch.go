package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/patch"
)

type PatchRepository interface {
	GetLatest(ctx context.Context, environment uuid.UUID) (patch.Patch, error)
	GetList(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error)
}

type PatchService interface {
	Get(ctx context.Context, environment uuid.UUID) (patch.Patch, error)
	GetList(ctx context.Context, environment uuid.UUID, limit int) ([]patch.Patch, error)
}
