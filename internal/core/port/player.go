package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/player"
)

type SocialPlayerRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (player.SocialPlayer, error)
}

type PlayerSearchService interface {
	Search(ctx context.Context, customer string) (player.SocialPlayer, error)
	// Get is Search for a caller that already holds the player ID - a client that played a
	// multiplayer game with them - rather than the customer ID a player shares.
	Get(ctx context.Context, id uuid.UUID) (player.SocialPlayer, error)
}
