package port

import (
	"context"
	"mevway/internal/core/domain/content"
)

type ProgressRepository interface {
	GetReleaseList(ctx context.Context) ([]content.GameFeatureRelease, error)
	GetProgressList(ctx context.Context) ([]content.GameFeature, error)
}

type ProgressService interface {
	ListRelease(ctx context.Context) ([]content.GameFeatureRelease, error)
	ListProgress(ctx context.Context) ([]content.GameFeature, error)
}
