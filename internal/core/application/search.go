package application

import (
	"context"
	"mevway/internal/core/domain/player"
	"mevway/internal/core/port"
)

type PlayerSearchService struct {
	identity port.IdentityRepository
	social   port.SocialPlayerRepository
}

func NewPlayerSearchService(identity port.IdentityRepository, social port.SocialPlayerRepository) *PlayerSearchService {
	return &PlayerSearchService{identity: identity, social: social}
}

func (s *PlayerSearchService) Search(ctx context.Context, customer string) (player.SocialPlayer, error) {

	identity, err := s.identity.IdentityFromCustomerID(ctx, customer)
	if err != nil {
		return player.SocialPlayer{}, err
	}

	social, err := s.social.GetByID(ctx, identity.PlayerID)
	if err != nil {
		return player.SocialPlayer{}, err
	}

	return social, err

}
