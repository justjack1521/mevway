package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
)

type PrivateAPIRouter struct {
	BaseAPIRouter
	BanUserHandler  handler.BanUserHandler
	UserRoleHandler handler.UserRoleHandler
}

func (a *PrivateAPIRouter) HandleAdminRoleAuthorise(ctx *gin.Context) {

	user, err := a.user(ctx)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	a.UserRoleHandler.Handle(ctx, handler.UserRole{
		UserID:   user,
		RoleName: "admin",
	})

}

func (a *PrivateAPIRouter) HandleBanUser(ctx *gin.Context) {

	request, err := resources.Binder[resources.BanUserRequest](ctx, resources.BanUserRequest{})

	if err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.BanUserHandler.Handle(ctx, handler.BanUser{
		UserID: request.SysUser,
	})
}

func (a *PrivateAPIRouter) ApplyRouterDecorations(router *gin.Engine) {

	private := router.Group("/private")

	private.Use(a.HandleTokenAuthorise)
	private.Use(a.HandleAdminRoleAuthorise)

	auth := private.Group("/session")
	auth.POST("/ban", a.HandleBanUser)

}
