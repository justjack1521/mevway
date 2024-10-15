package external

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protoadmin"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/application"
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
