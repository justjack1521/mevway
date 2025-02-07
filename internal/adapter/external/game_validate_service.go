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
			Active:     card.Active,
			CardNumber: int32(card.CardNumber),
			InShop:     card.InShop,
			BaseAbilityCard: &protomodel.BaseAbilityCard{
				SysId:           card.BaseCard.SysID.String(),
				Active:          card.BaseCard.Active,
				Name:            card.BaseCard.Name,
				SkillSeedOne:    card.BaseCard.SkillSeedOne.String(),
				SkillSeedTwo:    card.BaseCard.SkillSeedTwo.String(),
				SkillSeedSplit:  card.BaseCard.SkillSeedSplit,
				SeedFusionBoost: int32(card.BaseCard.SeedFusionBoost),
				Ability:         card.BaseCard.AbilityID.String(),
				Element:         card.BaseCard.Element.String(),
				Category:        card.BaseCard.Category,
				FastLearner:     card.BaseCard.FastLearner,
			},
			FusionExpOverride: int32(card.FusionEXPOverride),
			SaleGoldOverride:  int32(card.SaleGoldOverride),
			OverrideAbility:   card.OverrideAbilityID.String(),
		},
	}

	_, err := s.svc.ValidateAbilityCard(OutgoingContext(ctx), request)
	if err != nil {
		return err
	}

	return nil

}
