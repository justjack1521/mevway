package application

import (
	"context"
	"mevway/internal/core/domain/player"
	"mevway/internal/core/port"

	uuid "github.com/satori/go.uuid"
)

type RankQueryService struct {
	repository port.RankRepository
}

func NewRankQueryService(repository port.RankRepository) *RankQueryService {
	return &RankQueryService{repository: repository}
}

func (r *RankQueryService) ListTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error) {
	return r.repository.QueryRankingRange(ctx, code, uuid.Nil, 0, 10)
}

func (r *RankQueryService) ListPlayerRanking(ctx context.Context, code string, player uuid.UUID) (player.RankPlayer, error) {
	return r.repository.QueryPlayerRanking(ctx, code, player)
}

func (r *RankQueryService) ListRankingRange(ctx context.Context, code string, player uuid.UUID, start, stop int) ([]player.RankPlayer, error) {
	return r.repository.QueryRankingRange(ctx, code, player, start, stop)
}
