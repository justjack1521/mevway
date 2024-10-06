package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/user"
	"mevway/internal/decorator"
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

}
