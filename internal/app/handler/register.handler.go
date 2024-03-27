package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
	"mevway/internal/resources"
)

type RegisterUser struct {
	Username        string
	Password        string
	ConfirmPassword string
}

type RegisterUserHandler decorator.APIRouterHandler[RegisterUser]

type registerUserHandler struct {
	client services.AccessServiceClient
}

func NewRegisterUserHandler(clt services.AccessServiceClient) RegisterUserHandler {
	return registerUserHandler{
		client: clt,
	}
}

func (h registerUserHandler) Handle(ctx *gin.Context, query RegisterUser) {

	response, err := h.client.RegisterUser(ctx, &protoaccess.RegisterUserRequest{
		Username:        query.Username,
		Password:        query.Password,
		ConfirmPassword: query.ConfirmPassword,
	})

	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	user, err := uuid.FromString(response.UserId)
	if err != nil {
		httperr.InternalError(err, err.Error(), ctx)
		return
	}

	ctx.JSON(200, resources.UserRegisterResponse{SysUser: user})

}
