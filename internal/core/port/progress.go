package port

import (
	"context"
	"mevway/internal/core/domain/progress"
)

type ProgressRepository interface {
	GetProgressList(ctx context.Context) ([]progress.GameFeature, error)
}

type ProgressService interface {
	ListProgress(ctx context.Context) ([]progress.GameFeature, error)
}
