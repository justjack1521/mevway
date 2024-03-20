package ports

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"io/ioutil"
	"mevway/internal/app/handler"
	"mevway/internal/resources"
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
		UserID:   ctx.GetHeader("X-API-CLIENT"),
		RoleName: "alpha_tester",
	})
}

func (a *PublicAPIRouter) HandleSocket(ctx *gin.Context) {

	a.WebsocketHandle.Handle(ctx, handler.WebSocketQuery{
		ClientID: ctx.GetHeader("X-API-CLIENT"),
	})
}

func (a *PublicAPIRouter) HandleLoginUser(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}
	fmt.Println("Received request body:")
	fmt.Println(string(body))

	request, err := resources.Binder[resources.UserLoginRequest](ctx, resources.UserLoginRequest{})
	if err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.LoginUserHandle.Handle(ctx, handler.LoginUser{
		Username:   request.Username,
		Password:   request.Password,
		RememberMe: request.RememberMe,
	})

}

func (a *PublicAPIRouter) HandleRegisterUser(ctx *gin.Context) {

	request, err := resources.Binder[resources.UserRegisterRequest](ctx, resources.UserRegisterRequest{})

	if err != nil {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.RegisterUserHandle.Handle(ctx, handler.RegisterUser{
		Username:        request.Username,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	})

}

func (a *PublicAPIRouter) HandlePlayerSearch(ctx *gin.Context) {

	request, err := resources.Binder[resources.PlayerSearchRequest](ctx, resources.PlayerSearchRequest{})

	request.CustomerID = ctx.Param("customer_id")

	if err != nil || request.CustomerID == "" {
		httperr.BadRequest(err, "Bad request", ctx)
		return
	}

	a.PlayerSearchHandle.Handle(ctx, handler.PlayerSearch{CustomerID: request.CustomerID})

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
