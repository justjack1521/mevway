package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protosocial"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
	"mevway/internal/resources"
)

type PlayerSearch struct {
	UserID     uuid.UUID
	CustomerID string
}

type PlayerSearchHandler decorator.APIRouterHandler[PlayerSearch]

type playerSearchHandler struct {
	social services.MeviusSocialServiceClient
	cache  CustomerPlayerIDReadRepository
}

func NewPlayerSearchHandler(social services.MeviusSocialServiceClient, cache CustomerPlayerIDReadRepository) PlayerSearchHandler {
	return playerSearchHandler{social: social, cache: cache}
}

func (h playerSearchHandler) Handle(ctx *gin.Context, query PlayerSearch) {

	player, err := h.cache.Get(ctx, query.CustomerID)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	search, err := h.social.PlayerSearch(ctx, &protosocial.PlayerSearchRequest{PlayerId: player.String()})
	if err != nil || search.PlayerInfo == nil || search.PlayerInfo.PlayerInfo.PlayerId == uuid.Nil.String() {
		httperr.BadRequest(err, "player not found", ctx)
		return
	}

	var response = &resources.PlayerSearchResponse{
		PlayerID:      uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.PlayerId),
		PlayerName:    search.PlayerInfo.PlayerInfo.PlayerName,
		PlayerLevel:   int(search.PlayerInfo.PlayerInfo.PlayerLevel),
		PlayerComment: search.PlayerInfo.PlayerInfo.PlayerComment,
		CompanionID:   uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.CompanionId),
		JobCardID:     uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.CompanionId),
		SubJobIndex:   int(search.PlayerInfo.PlayerInfo.SubJobIndex),
	}

	if search.PlayerInfo.PlayerInfo.RentalCard != nil {
		response.RentalCard = &resources.PlayerSearchResponseRentalCard{
			AbilityCardID:    uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.RentalCard.AbilityCardId),
			AbilityCardLevel: int(search.PlayerInfo.PlayerInfo.RentalCard.AbilityCardLevel),
			AbilityLevel:     int(search.PlayerInfo.PlayerInfo.RentalCard.AbilityLevel),
			ExtraSkillUnlock: int(search.PlayerInfo.PlayerInfo.RentalCard.ExtraSkillUnlock),
			OverBoostLevel:   int(search.PlayerInfo.PlayerInfo.RentalCard.OverBoostLevel),
		}
	}

	ctx.JSON(200, response)

}
