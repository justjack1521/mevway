package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
)

const (
	UserIDContextKey   string = "UserIDContextKey"
	PlayerIDContextKey string = "PlayerIDContextKey"
	UserEnvironmentKey string = "UserEnvironmentKey"
)

type TokenAuthorise struct {
	SessionID uuid.UUID
	Bearer    string
	DeviceID  string
}

type TokenAuthoriseHandler decorator.APIRouterHandler[TokenAuthorise]

type tokenAuthoriseHandler struct {
	client services.AccessServiceClient
}

func NewTokenHandler(clt services.AccessServiceClient) TokenAuthoriseHandler {
	return tokenAuthoriseHandler{
		client: clt,
	}
}

func (h tokenAuthoriseHandler) Handle(ctx *gin.Context, query TokenAuthorise) {

	response, err := h.client.AuthenticateToken(ctx, &protoaccess.AuthenticateTokenRequest{
		SessionId: query.SessionID.String(),
		Bearer:    query.Bearer,
		DeviceId:  query.DeviceID,
	})

	if err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

	ctx.Set(UserIDContextKey, response.UserId)
	ctx.Set(PlayerIDContextKey, response.PlayerId)
	ctx.Set(UserEnvironmentKey, response.Environment)

}
