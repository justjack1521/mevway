package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
	"mevway/internal/domain/user"
	"mevway/internal/resources"
)

type RegisterUser struct {
	Username        string
	Password        string
	ConfirmPassword string
}

type RegisterUserHandler decorator.APIRouterHandler[RegisterUser]

type registerUserHandler struct {
	client RegistrationClient
}

type RegistrationClient interface {
	Register(ctx context.Context, user user.User) (uuid.UUID, error)
}

func NewRegisterUserHandler(clt RegistrationClient) RegisterUserHandler {
	return registerUserHandler{
		client: clt,
	}
}

func (h registerUserHandler) Handle(ctx *gin.Context, query RegisterUser) {

	response, err := h.client.Register(ctx, user.NewUser(query.Username, query.Password))

	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	ctx.JSON(200, resources.UserRegisterResponse{SysUser: response})

}
