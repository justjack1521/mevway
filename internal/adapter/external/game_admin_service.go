package external

import (
	"context"
	"mevway/internal/core/domain/game"

	"github.com/justjack1521/mevium/pkg/genproto/protoadmin"
	"github.com/justjack1521/mevium/pkg/genproto/protogame"
	"github.com/justjack1521/mevium/pkg/genproto/protomodel"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	uuid "github.com/satori/go.uuid"
)

type GameAdminService struct {
	svc services.MeviusAdminServiceClient
}

func (s *GameAdminService) CreateBaseCard(ctx context.Context, card game.BaseCard) error {

	var request = &protoadmin.CreateBaseCardRequest{
		Card: &protomodel.BaseAbilityCard{
			SysId:               card.SysID.String(),
			Active:              card.Active,
			Name:                card.Name,
			FiendCard:           false,
			SkillSeedOne:        card.SkillSeedOne.String(),
			SkillSeedTwo:        card.SkillSeedTwo.String(),
			SkillSeedSplit:      card.SkillSeedSplit,
			SeedFusionBoost:     int32(card.SeedFusionBoost),
			Ability:             card.AbilityID.String(),
			FastLearner:         card.FastLearner,
			ExpFusionMultiplier: card.EXPFusionMultiplier,
		},
	}

	_, err := s.svc.CreateBaseCard(ctx, request)
	if err != nil {
		return err
	}

	return nil

}

func NewGameAdminService(svc services.MeviusAdminServiceClient) *GameAdminService {
	return &GameAdminService{svc: svc}
}

func (s *GameAdminService) GrantItem(ctx context.Context, player uuid.UUID, item uuid.UUID, quantity int) error {

	var request = &protoadmin.GrantItemRequest{
		PlayerId: player.String(),
		ItemId:   item.String(),
		Quantity: int32(quantity),
	}

	_, err := s.svc.GrantItem(OutgoingContext(ctx), request)
	if err != nil {
		return err
	}

	return nil

}

func (s *GameAdminService) CreateBaseJob(ctx context.Context, job game.BaseJob) (bool, error) {

	var request = &protoadmin.CreateBaseJobCardRequest{
		SysId:     job.SysID.String(),
		Active:    job.Active,
		JobNumber: job.Number,
		JobName:   job.Name,
		TypeId:    job.TypeID.String(),
	}

	response, err := s.svc.CreateBaseJobCard(OutgoingContext(ctx), request)
	if err != nil {
		return false, err
	}

	return response.Created, nil

}

func (s *GameAdminService) CreateSkillPanel(ctx context.Context, job uuid.UUID, page int, panel game.SkillPanel) (bool, error) {

	var request = &protoadmin.CreateSkillPanelRequest{
		BaseJobId:      job.String(),
		PageIndex:      int32(page),
		DefinitionType: panel.DefinitionType,
		Index:          int32(panel.Index),
		ReferenceId:    panel.ReferenceID.String(),
		Value:          int32(panel.Value),
		CostItems:      make([]*protogame.ProtoItemValuePair, len(panel.CostItems)),
	}

	for index, value := range panel.CostItems {
		request.CostItems[index] = &protogame.ProtoItemValuePair{
			ItemId: value.ItemID.String(),
			Value:  int32(value.Value),
		}
	}

	response, err := s.svc.CreateSkillPanel(OutgoingContext(ctx), request)
	if err != nil {
		return false, err
	}

	return response.Created, nil

}

func (s *GameAdminService) CreateAugmentMaterials(ctx context.Context, id uuid.UUID, materials []game.AugmentMaterial) error {

	var request = &protoadmin.CreateAugmentMaterialRequest{
		AbilityCardId: id.String(),
		Materials:     make(map[string]int32, 0),
	}

	for _, value := range materials {
		request.Materials[value.SysID.String()] = int32(value.Quantity)
	}

	_, err := s.svc.CreateAugmentMaterials(OutgoingContext(ctx), request)
	if err != nil {
		return err
	}

	return nil

}
