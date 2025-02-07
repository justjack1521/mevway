package external

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protomodel"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"mevway/internal/core/domain/game"
)

type GameValidateService struct {
	svc services.MeviusModelServiceClient
}

func NewGameValidateService(svc services.MeviusModelServiceClient) *GameValidateService {
	return &GameValidateService{svc: svc}
}

func (s *GameValidateService) ValidateAbilityCard(ctx context.Context, card game.AbilityCard) error {

	var request = &protomodel.ValidateAbilityCardRequest{
		AbilityCard: &protomodel.AbilityCard{
			SysId:      card.SysID.String(),
			Active:     false,
			CardNumber: 0,
			InShop:     false,
			BaseAbilityCard: &protomodel.BaseAbilityCard{
				SysId:           card.BaseCard.SysID.String(),
				Active:          false,
				Name:            card.BaseCard.Name,
				FiendCard:       false,
				SkillSeedOne:    "",
				SkillSeedTwo:    "",
				SkillSeedSplit:  "",
				SeedFusionBoost: 0,
				Ability:         card.BaseCard.AbilityID.String(),
				Element:         "",
				Category:        "",
				FastLearner:     false,
			},
			AugmentConfig:     nil,
			FusionExpOverride: 0,
			SaleGoldOverride:  0,
			OverrideAbility:   "",
		},
	}

	_, err := s.svc.ValidateAbilityCard(ctx, request)
	if err != nil {
		return err
	}

	return nil

}
