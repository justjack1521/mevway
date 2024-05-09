package handler

import (
	"context"
	uuid "github.com/satori/go.uuid"
)

type CustomerPlayerIDReadRepository interface {
	Get(ctx context.Context, customer string) (uuid.UUID, error)
}

type CustomerPlayerIDWriteRepository interface {
	Add(ctx context.Context, customer string, player uuid.UUID) error
}

type CustomerPlayerIDRepository interface {
	CustomerPlayerIDReadRepository
	CustomerPlayerIDWriteRepository
}
