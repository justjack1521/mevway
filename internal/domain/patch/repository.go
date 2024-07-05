package patch

import "context"

type ReadRepository interface {
	Get(ctx context.Context, limit int) ([]Patch, error)
}

type Repository interface {
	ReadRepository
}
