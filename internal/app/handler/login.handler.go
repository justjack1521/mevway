package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
	"mevway/internal/resources"
)

type LoginUser struct {
	Username string
	Password string
	DeviceID string
}

type CustomerIDCache interface {
	GetUserIDFromCustomerID(customer string) (uuid.UUID, error)
	AddCustomerIDForUser(customer string, user uuid.UUID) error
}

type LoginUserHandler decorator.APIRouterHandler[LoginUser]

type loginUserHandler struct {
	client services.AccessServiceClient
	cache  CustomerIDCache
}

func NewLoginHandler(clt services.AccessServiceClient, cache CustomerIDCache) LoginUserHandler {
	return loginUserHandler{
		client: clt,
		cache:  cache,
	}
}

func (h loginUserHandler) Handle(ctx *gin.Context, query LoginUser) {

	login, err := h.client.LoginUser(ctx, &protoaccess.LoginUserRequest{
		Username: query.Username,
		Password: query.Password,
		DeviceId: query.DeviceID,
	})
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	//TODO Remove when opening build
	role, err := h.client.UserHasRole(ctx, &protoaccess.UserHasRoleRequest{
		UserId: login.UserId,
		Role:   "alpha_tester",
	})
	if err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

	if role.HasRole == false {
		err := errors.New("unauthorised")
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

	ctx.JSON(200, resources.UserLoginResponse{
		SessionID:    login.SessionId,
		CustomerID:   login.CustomerId,
		AccessToken:  login.AccessToken,
		RefreshToken: login.RefreshToken,
	})

}
