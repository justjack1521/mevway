package cache

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/app/handler"
)

type CustomerPlayerIDRepository struct {
	source handler.CustomerPlayerIDReadRepository
	cache  handler.CustomerPlayerIDRepository
}

func NewCustomerPlayerIDRepository(source handler.CustomerPlayerIDReadRepository, cache handler.CustomerPlayerIDRepository) *CustomerPlayerIDRepository {
	return &CustomerPlayerIDRepository{source: source, cache: cache}
}

func (r *CustomerPlayerIDRepository) Get(ctx context.Context, customer string) (uuid.UUID, error) {
	cached, err := r.cache.Get(ctx, customer)
	if err == nil {
		return cached, nil
	}
	actual, err := r.source.Get(ctx, customer)
	if err != nil {
		return uuid.Nil, err
	}
	r.cache.Add(ctx, customer, actual)
	return actual, nil
}
