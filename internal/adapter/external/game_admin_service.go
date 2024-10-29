package external

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protoadmin"
	"github.com/justjack1521/mevium/pkg/genproto/protogame"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/game"
)

type GameAdminService struct {
	svc services.MeviusAdminServiceClient
}

func NewGameAdminService(svc services.MeviusAdminServiceClient) *GameAdminService {
	return &GameAdminService{svc: svc}
}

func (s *GameAdminService) GrantItem(ctx context.Context, player uuid.UUID, item uuid.UUID, quantity int) error {

	var md = application.MetadataFromContext(ctx)

	var request = &protoadmin.GrantItemRequest{
		PlayerId: player.String(),
		ItemId:   item.String(),
		Quantity: int32(quantity),
	}

	_, err := s.svc.GrantItem(mevrpc.NewOutgoingContext(ctx, md.UserID, md.PlayerID), request)
	if err != nil {
		return err
	}

	return nil

}

func (s *GameAdminService) CreateSkillPanel(ctx context.Context, job uuid.UUID, page int, panel game.SkillPanel) error {

	var md = application.MetadataFromContext(ctx)

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

	_, err := s.svc.CreateSkillPanel(mevrpc.NewOutgoingContext(ctx, md.UserID, md.PlayerID), request)
	if err != nil {
		return err
	}

	return nil

}
