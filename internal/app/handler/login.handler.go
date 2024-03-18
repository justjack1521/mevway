package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/status"
	"mevway/internal/decorator"
	"mevway/internal/resources"
)

type LoginUser struct {
	Username   string
	Password   string
	RememberMe bool
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

	response, err := h.client.LoginUser(ctx, &protoaccess.LoginUserRequest{
		Username:   query.Username,
		Password:   query.Password,
		RememberMe: query.RememberMe,
	})

	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			httperr.BadRequest(err, st.Message(), ctx)
		}
		return
	}

	user, err := uuid.FromString(response.Header.ClientId)
	if err != nil {
		httperr.InternalError(err, err.Error(), ctx)
		return
	}

	//TODO Remove when opening build
	_, err = h.client.UserHasRole(ctx, &protoaccess.UserHasRoleRequest{
		Header: &protocommon.RequestHeader{ClientId: response.Header.ClientId},
		Role:   "alpha_tester",
	})

	if err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
	}

	if err := h.cache.AddCustomerIDForUser(response.CustomerId, user); err != nil {
		httperr.InternalError(err, err.Error(), ctx)
		return
	}

	ctx.JSON(200, resources.UserLoginResponse{
		SysUser:      response.Header.ClientId,
		CustomerID:   response.CustomerId,
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		RememberMe:   response.RememberMe,
	})

}
