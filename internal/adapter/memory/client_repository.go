package memory

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"mevway/internal/adapter/memory/dto"
	"mevway/internal/core/domain/socket"
	"strings"
	"time"
)

const (
	connectedClientsKey = "connected_clients"
	connectedClientsTTL = time.Hour * 8
)

type ClientConnectionRepository struct {
	client *redis.Client
}

func NewClientRepository(client *redis.Client) *ClientConnectionRepository {
	return &ClientConnectionRepository{client: client}
}

func (c *ClientConnectionRepository) Add(ctx context.Context, client socket.Client) error {

	var key = c.key(client)

	var hash = dto.SocketClientRedis{
		SessionID: client.Session.String(),
		UserID:    client.UserID.String(),
		PlayerID:  client.PlayerID.String(),
	}

	if err := c.client.HSet(ctx, key, hash.ToMapStringInterface()).Err(); err != nil {
		return err
	}

	if err := c.client.Expire(ctx, key, connectedClientsTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (c *ClientConnectionRepository) Remove(ctx context.Context, client socket.Client) error {
	if err := c.client.Del(ctx, c.key(client)).Err(); err != nil {
		return err
	}
	return nil
}

func (c *ClientConnectionRepository) List(ctx context.Context) ([]socket.Client, error) {
	result, _, err := c.client.HScan(ctx, connectedClientsKey, 0, "", 10).Result()
	if err != nil {
		return nil, err
	}
	fmt.Print(result)
	return nil, nil
}

func (c *ClientConnectionRepository) key(client socket.Client) string {
	return strings.Join([]string{serviceKey, connectedClientsKey, client.Session.String()}, ":")
}
