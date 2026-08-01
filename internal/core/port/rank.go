package port

import (
	"context"
	"mevway/internal/core/domain/player"

	uuid "github.com/satori/go.uuid"
)

type RankRepository interface {
	QueryTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error)
	QueryPlayerRanking(ctx context.Context, code string, player uuid.UUID) (player.RankPlayer, error)
	QueryRankingRange(ctx context.Context, code string, player uuid.UUID, start, stop int) ([]player.RankPlayer, error)
}

type RankService interface {
	ListTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error)
	ListPlayerRanking(ctx context.Context, code string, player uuid.UUID) (player.RankPlayer, error)
	ListRankingRange(ctx context.Context, code string, player uuid.UUID, start, stop int) ([]player.RankPlayer, error)
}
