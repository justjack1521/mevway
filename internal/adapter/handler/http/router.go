package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/middleware"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	authHandler *AuthenticationHandler,
	statusHandler *StatusHandler,
	patchHandler *PatchHandler,
	socketHandler *SocketHandler,
	searchHandler *SearchHandler,
	middle ...gin.HandlerFunc,
) (*Router, error) {

	var router = gin.New()
	router.HandleMethodNotAllowed = false

	router.Use(middleware.CORSMiddleware())
	router.Use(middle...)

	var publicGroup = router.Group("/public")

	var socketGroup = publicGroup.Group("/socket", authHandler.TokenAuthorise)
	{
		socketGroup.GET("/join", socketHandler.Join)
		socketGroup.GET("/list", middleware.AdminRoleMiddleware(), socketHandler.List)
	}

	var authGroup = publicGroup.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/register", authHandler.Register)
	}

	var systemGroup = publicGroup.Group("/system")
	{
		systemGroup.GET("/status", statusHandler.Get)
	}

	var patch = publicGroup.Group("/patch", authHandler.TokenAuthorise)
	{
		patch.GET("/recent", patchHandler.Recent)
		patch.GET("/list", patchHandler.List)
	}

	var player = publicGroup.Group("/player", authHandler.TokenAuthorise)
	{
		player.GET("/search/:customer_id", searchHandler.Search)
	}

	return &Router{router}, nil

}

func (r *Router) Serve(addr string) error {
	return r.Run(addr)
}
