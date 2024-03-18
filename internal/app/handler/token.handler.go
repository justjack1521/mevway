package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
)

type TokenAuthorise struct {
	UserID string
	Bearer string
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

	resp, err := h.client.AuthToken(ctx, &protoaccess.AuthTokenRequest{
		Header: &protocommon.RequestHeader{
			ClientId: query.UserID,
		},
		Bearer: query.Bearer,
	})

	if err != nil {
		httperr.ResponseError(resp.Header, ctx)
		return
	}

}
