package port

import (
	"context"
	"mevway/internal/core/domain/player"
)

type RankRepository interface {
	QueryTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error)
}

type RankService interface {
	ListTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error)
}
