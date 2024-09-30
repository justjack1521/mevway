package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
	"mevway/internal/domain/auth"
	"mevway/internal/resources"
)

type LoginUser struct {
	Username   string
	Password   string
	DeviceID   string
	RememberMe bool
}

type LoginClient interface {
	Login(ctx context.Context, request auth.LoginRequest) (auth.LoginResponse, error)
}

type LoginUserHandler decorator.APIRouterHandler[LoginUser]

type loginUserHandler struct {
	client LoginClient
}

func NewLoginHandler(clt LoginClient) LoginUserHandler {
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
		IDToken:      login.IDToken,
		AccessToken:  login.AccessToken,
		RefreshToken: login.RefreshToken,
		//RememberToken: login.RememberToken,
	})

}
