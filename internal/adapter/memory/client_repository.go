package memory

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/memory/dto"
	"mevway/internal/core/domain/socket"
	"strings"
	"time"
)

const (
	connectedClientInfoKey = "connected_client_info"
	connectedClientsTTL    = time.Hour * 8
)

type ClientConnectionRepository struct {
	client *redis.Client
}

func NewClientRepository(client *redis.Client) *ClientConnectionRepository {
	return &ClientConnectionRepository{client: client}
}

func (c *ClientConnectionRepository) Add(ctx context.Context, client socket.Client) error {

	var key = c.clientKey(client)

	var hash = dto.SocketClientRedis{
		SessionID: client.Session.String(),
		UserID:    client.UserID.String(),
		PlayerID:  client.PlayerID.String(),
		PatchID:   client.PatchID.String(),
	}

	if err := c.client.HSet(ctx, key, hash.ToMapStringInterface()).Err(); err != nil {
		return err
	}

	if err := c.client.Expire(ctx, key, connectedClientsTTL).Err(); err != nil {
		return err
	}

	return nil
}

func (c *ClientConnectionRepository) RemoveAll(ctx context.Context) error {
	iter := c.client.Scan(ctx, 0, fmt.Sprintf("%s*", c.clientKey(socket.Client{})), 0).Iterator()
	for iter.Next(ctx) {
		if err := c.client.Del(ctx, iter.Val()).Err(); err != nil {
			return err
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}

func (c *ClientConnectionRepository) Remove(ctx context.Context, client socket.Client) error {
	if err := c.client.Del(ctx, c.clientKey(client)).Err(); err != nil {
		return err
	}
	return nil
}

func (c *ClientConnectionRepository) List(ctx context.Context) ([]socket.Client, error) {
	result, _, err := c.client.HScan(ctx, connectedClientInfoKey, 0, "", 10).Result()
	if err != nil {
		return nil, err
	}
	fmt.Print(result)
	return nil, nil
}

func (c *ClientConnectionRepository) clientKey(client socket.Client) string {

	var id = ""

	if uuid.Equal(client.Session, uuid.Nil) == false {
		id = client.Session.String()
	}

	return strings.Join([]string{serviceKey, connectedClientInfoKey, id}, ":")
}
