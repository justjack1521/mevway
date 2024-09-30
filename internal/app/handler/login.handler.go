package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
	"mevway/internal/domain/auth"
	"mevway/internal/ports"
	"mevway/internal/resources"
)

type LoginUser struct {
	Username   string
	Password   string
	DeviceID   string
	RememberMe bool
}

type LoginUserHandler decorator.APIRouterHandler[LoginUser]

type loginUserHandler struct {
	client ports.AuthenticationClient
}

func NewLoginHandler(clt ports.AuthenticationClient) LoginUserHandler {
	return loginUserHandler{
		client: clt,
	}
}

func (h loginUserHandler) Handle(ctx *gin.Context, query LoginUser) {

	login, err := h.client.Login(ctx, auth.LoginRequest{
		Username: query.Username,
		Password: query.Password,
	})

	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	ctx.JSON(200, resources.UserLoginResponse{
		//SessionID:     login.SessionId,
		//CustomerID:    login.CustomerId,
		AccessToken:  login.AccessToken,
		RefreshToken: login.RefreshToken,
		//RememberToken: login.RememberToken,
	})

}
