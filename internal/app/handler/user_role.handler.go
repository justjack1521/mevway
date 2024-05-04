package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/decorator"
)

type UserRole struct {
	UserID   uuid.UUID
	RoleName string
}

type UserRoleHandler decorator.APIRouterHandler[UserRole]

type userRoleHandler struct {
	client services.AccessServiceClient
}

func NewUserRoleHandler(clt services.AccessServiceClient) UserRoleHandler {
	return userRoleHandler{
		client: clt,
	}
}

func (h userRoleHandler) Handle(ctx *gin.Context, query UserRole) {

	response, err := h.client.UserHasRole(ctx, &protoaccess.UserHasRoleRequest{
		UserId: query.UserID.String(),
		Role:   query.RoleName,
	})

	if err != nil {
		httperr.UnauthorisedError(err, err.Error(), ctx)
		return
	}

	if response.HasRole == false {
		httperr.UnauthorisedError(errors.New("unauthorised"), "unauthorised", ctx)
		return
	}

}
