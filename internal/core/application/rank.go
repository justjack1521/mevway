package application

import (
	"context"
	"mevway/internal/core/domain/player"
	"mevway/internal/core/port"
)

type RankQueryService struct {
	repository port.RankRepository
}

func (r *RankQueryService) ListTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error) {
	return r.repository.QueryTopRankings(ctx, code)
}

func NewRankQueryService(repository port.RankRepository) *RankQueryService {
	return &RankQueryService{repository: repository}
}
