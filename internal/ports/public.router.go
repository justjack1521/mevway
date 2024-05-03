package ports

import (
	"github.com/gin-gonic/gin"
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
	a.UserRoleHandler.Handle(ctx, handler.UserRole{
		UserID:   a.session(ctx),
		RoleName: "alpha_tester",
	})
}

func (a *PublicAPIRouter) ApplyRouterDecorations(router *gin.Engine) {

	pub := router.Group("/public")

	pub.GET("/ws", a.HandleTokenAuthorise, a.HandlerAlphaTesterAuthorise, a.HandleSocket)

	auth := pub.Group("/auth")
	auth.POST("/login", a.HandleLoginUser)
	auth.POST("/register", a.HandleRegisterUser)

	search := pub.Group("/player_search", a.HandlerAlphaTesterAuthorise, a.HandleTokenAuthorise)
	search.GET("/:customer_id", a.HandlePlayerSearch)

}
