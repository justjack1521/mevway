package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
	"mevway/internal/resources"
)

type RememberUser struct {
	Token    string
	DeviceID string
}

type RememberUserHandler decorator.APIRouterHandler[RememberUser]

type rememberUserHandler struct {
	client services.AccessServiceClient
}

func NewRememberUserHandler(clt services.AccessServiceClient) RememberUserHandler {
	return rememberUserHandler{
		client: clt,
	}
}

func (h rememberUserHandler) Handle(ctx *gin.Context, query RememberUser) {

	rmb, err := h.client.RememberUser(ctx, &protoaccess.RememberUserRequest{
		RememberToken: query.Token,
		DeviceId:      query.DeviceID,
	})
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	ctx.JSON(200, resources.RememberUserResponse{
		SessionID:    rmb.SessionId,
		CustomerID:   rmb.CustomerId,
		Username:     rmb.Username,
		AccessToken:  rmb.AccessToken,
		RefreshToken: rmb.RefreshToken,
	})

}
