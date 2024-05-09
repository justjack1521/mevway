package memory

import (
	"context"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

var (
	customerPlayerCacheKey = "customer_player_cache"
)

type CustomerPlayerIDRepository struct {
	client *redis.Client
}

func NewCustomerPlayerCache(client *redis.Client) *CustomerPlayerIDRepository {
	return &CustomerPlayerIDRepository{client: client}
}

func (r *CustomerPlayerIDRepository) Get(ctx context.Context, customer string) (uuid.UUID, error) {
	result, err := r.client.Get(ctx, r.key(customer)).Result()
	if err != nil {
		return uuid.Nil, err
	}
	value, err := uuid.FromString(result)
	if err != nil {
		return uuid.Nil, err
	}
	return value, nil
}

func (r *CustomerPlayerIDRepository) Add(ctx context.Context, customer string, player uuid.UUID) error {
	if err := r.client.Set(ctx, r.key(customer), player.String(), time.Minute*60).Err(); err != nil {
		return err
	}
	return nil
}

func (r *CustomerPlayerIDRepository) key(customer string) string {
	return strings.Join([]string{serviceKey, customerPlayerCacheKey, customer}, ":")
}
