package rpc

import (
	"context"
	"errors"
	"fmt"
	"mevway/internal/core/domain/socket"
)

var (
	errFailedRoutingClientRequest = func(err error) error {
		return fmt.Errorf("failed to route client request: %w", err)
	}
	errMalformedRequest = errors.New("malformed request")
)

const (
	GameClientRouteKey      = 100
	SocialClientRouteKey    = 200
	RankingClientRouteKey   = 300
	ChallengeClientRouteKey = 400
	MultiClientRouteKey     = 500
)

type handler func(ctx context.Context, bytes []byte) (socket.Response, error)
