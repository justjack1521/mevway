package patch

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type ReadRepository interface {
	Current(ctx context.Context, environment uuid.UUID) (Patch, error)
	Get(ctx context.Context, environment uuid.UUID, limit int) ([]Patch, error)
}

type Repository interface {
	ReadRepository
}
