package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/app/handler"
)

type PublicAPIRouter struct {
	BaseAPIRouter
	LoginUserHandle    handler.LoginUserHandler
	RegisterUserHandle handler.RegisterUserHandler
	WebsocketHandle    handler.WebSocketHandler
	PlayerSearchHandle handler.PlayerSearchHandler
	UserRoleHandler    handler.UserRoleHandler
}

func (a *PublicAPIRouter) HandlerAlphaTesterAuthorise(ctx *gin.Context) {

	user, err := a.user(ctx)
	if err != nil {
		httperr.BadRequest(err, err.Error(), ctx)
		return
	}

	a.UserRoleHandler.Handle(ctx, handler.UserRole{
		UserID:   user,
		RoleName: "alpha_tester",
	})
}

func (a *PublicAPIRouter) ApplyRouterDecorations(router *gin.Engine) {

	pub := router.Group("/public")

	pub.GET("/ws", a.HandleTokenAuthorise, a.HandleSocket)

	auth := pub.Group("/auth")
	auth.POST("/login", a.HandleLoginUser)
	auth.POST("/register", a.HandleRegisterUser)

	search := pub.Group("/player_search", a.HandleTokenAuthorise)
	search.GET("/:customer_id", a.HandleTokenAuthorise, a.HandlePlayerSearch)

}
