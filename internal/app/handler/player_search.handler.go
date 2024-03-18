package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/justjack1521/mevium/pkg/genproto/protosocial"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"mevway/internal/decorator"
	"mevway/internal/resources"
)

type PlayerSearch struct {
	UserID     string
	CustomerID string
}

type PlayerSearchHandler decorator.APIRouterHandler[PlayerSearch]

type playerSearchHandler struct {
	access services.AccessServiceClient
	social services.MeviusSocialServiceClient
	cache  CustomerIDCache
}

func NewPlayerSearchHandler(access services.AccessServiceClient, social services.MeviusSocialServiceClient, cache CustomerIDCache) PlayerSearchHandler {
	return playerSearchHandler{access: access, social: social, cache: cache}
}

func (h playerSearchHandler) Handle(ctx *gin.Context, query PlayerSearch) {

	md := metadata.New(map[string]string{"X-API-CLIENT": ctx.GetHeader("X-API-CLIENT")})
	out := metadata.NewOutgoingContext(ctx, md)
	user, err := h.cache.GetUserIDFromCustomerID(query.CustomerID)

	if err != nil {
		httperr.InternalError(err, err.Error(), ctx)
	}

	if user == uuid.Nil {
		result, err := h.access.CustomerSearch(out, &protoaccess.CustomerSearchRequest{
			Header:     &protocommon.RequestHeader{ClientId: query.UserID},
			CustomerId: query.CustomerID,
		})
		if err != nil {
			st, ok := status.FromError(err)
			if ok {
				httperr.BadRequest(err, st.Message(), ctx)
				return
			}
			httperr.BadRequest(err, err.Error(), ctx)
			return
		}
		user, err = uuid.FromString(result.UserId)
		if err != nil {
			httperr.BadRequest(err, err.Error(), ctx)
			return
		}
	}

	_ = h.cache.AddCustomerIDForUser(query.CustomerID, user)

	search, err := h.social.PlayerSearch(out, &protosocial.PlayerSearchRequest{UserId: user.String()})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			httperr.BadRequest(err, st.Message(), ctx)
			return
		}
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	if search.PlayerInfo == nil {
		httperr.BadRequest(err, "player not found", ctx)
		return
	}

	id, err := uuid.FromString(search.PlayerInfo.PlayerId)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	var response = &resources.PlayerSearchResponse{
		PlayerID:      id,
		PlayerName:    search.PlayerInfo.PlayerName,
		PlayerLevel:   int(search.PlayerInfo.PlayerLevel),
		PlayerComment: search.PlayerInfo.PlayerComment,
		CompanionID:   uuid.FromStringOrNil(search.PlayerInfo.CompanionId),
		JobCardID:     uuid.FromStringOrNil(search.PlayerInfo.JobCardId),
		SubJobIndex:   int(search.PlayerInfo.SubJobIndex),
	}

	if search.PlayerInfo.RentalCard != nil {
		response.RentalCard = &resources.PlayerSearchResponseRentalCard{
			AbilityCardID:    uuid.FromStringOrNil(search.PlayerInfo.RentalCard.AbilityCardId),
			AbilityCardLevel: int(search.PlayerInfo.RentalCard.AbilityCardLevel),
			AbilityLevel:     int(search.PlayerInfo.RentalCard.AbilityLevel),
			ExtraSkillUnlock: int(search.PlayerInfo.RentalCard.ExtraSkillUnlock),
			OverBoostLevel:   int(search.PlayerInfo.RentalCard.OverBoostLevel),
		}
	}

	ctx.JSON(200, response)

}
