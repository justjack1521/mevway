package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
)

type TokenAuthorise struct {
	UserID   string
	Bearer   string
	DeviceID string
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

	_, err := h.client.AuthenticateToken(ctx, &protoaccess.AuthenticateTokenRequest{
		UserId:   query.UserID,
		Bearer:   query.Bearer,
		DeviceId: query.DeviceID,
	})

	if err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

}
