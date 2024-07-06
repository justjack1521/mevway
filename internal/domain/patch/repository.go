package patch

import "context"

type ReadRepository interface {
	Current(ctx context.Context) (Patch, error)
	Get(ctx context.Context, limit int) ([]Patch, error)
}

type Repository interface {
	ReadRepository
}
