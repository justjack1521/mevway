package memory

import (
	"context"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

const (
	clientRequestKey = "client_request"
)

type RequestMemoryRepository struct {
	client *redis.Client
}

func NewRequestMemoryRepository(client *redis.Client) *RequestMemoryRepository {
	return &RequestMemoryRepository{client: client}
}

func (r *RequestMemoryRepository) Create(ctx context.Context, player uuid.UUID, bytes []byte) error {
	if err := r.client.Set(ctx, r.key(player), bytes, time.Hour*24).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RequestMemoryRepository) key(player uuid.UUID) string {
	return strings.Join([]string{
		serviceKey,
		clientRequestKey,
		player.String(),
		time.Now().String(),
	}, ":")
}
