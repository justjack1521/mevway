package memory

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"mevway/internal/adapter/memory/dto"
	"mevway/internal/domain/socket"
	"strings"
)

const (
	connectedClientsKey = "connected_clients"
)

type ClientRepository struct {
	client *redis.Client
}

func NewClientRepository(client *redis.Client) *ClientRepository {
	return &ClientRepository{client: client}
}

func (c *ClientRepository) Add(ctx context.Context, client socket.Client) error {
	var hash = dto.SocketClientRedis{
		UserID:   client.UserID.String(),
		PlayerID: client.PlayerID.String(),
	}
	if err := c.client.HMSet(ctx, c.key(client), hash).Err(); err != nil {
		return err
	}
	return nil
}

func (c *ClientRepository) Remove(ctx context.Context, client socket.Client) error {
	if err := c.client.Del(ctx, c.key(client)).Err(); err != nil {
		return err
	}
	return nil
}

func (c *ClientRepository) List(ctx context.Context) ([]socket.Client, error) {
	result, _, err := c.client.HScan(ctx, connectedClientsKey, 0, "", 10).Result()
	if err != nil {
		return nil, err
	}
	fmt.Print(result)
	return nil, nil
}

func (c *ClientRepository) key(client socket.Client) string {
	return strings.Join([]string{serviceKey, connectedClientsKey, client.Session.String()}, ":")
}
