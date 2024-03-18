package web

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

var (
	ErrFailedRoutingClientRequest = func(err error) error {
		return fmt.Errorf("failed to route client request: %w", err)
	}
	ErrSendingClientResponse = func(id uuid.UUID, err error) error {
		return fmt.Errorf("error sending client %v response: %w", id, err)
	}
	ErrMalformedRequest = errors.New("malformed request")
)

type RoutingKey int

const (
	GameClientRouteKey      = 100
	SocialClientRouteKey    = 200
	RankingClientRouteKey   = 300
	ChallengeClientRouteKey = 400
	MultiClientRouteKey     = 500
)

type ServiceClientRouter interface {
	Route(ctx *ClientContext, operation int, bytes []byte) (ClientResponse, error)
}

type ServiceClientRouteHandler func(ctx *ClientContext, bytes []byte) (ClientResponse, error)
