package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/genproto/protoaccess"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
)

type UserRole struct {
	UserID   string
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
		Header: &protocommon.RequestHeader{ClientId: query.UserID},
		Role:   query.RoleName,
	})

	if err != nil {
		if response == nil || response.Header == nil {
			httperr.UnauthorisedError(err, err.Error(), ctx)
			return
		}
		httperr.ResponseError(response.Header, ctx)
		return
	}

}
