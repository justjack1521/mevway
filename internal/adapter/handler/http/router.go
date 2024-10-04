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
	middle ...gin.HandlerFunc,
) (*Router, error) {

	var router = gin.New()
	router.HandleMethodNotAllowed = false

	router.Use(middleware.CORSMiddleware())
	router.Use(middle...)

	var publicGroup = router.Group("/public")

	publicGroup.GET("/ws", socketHandler.Handle)

	var authGroup = publicGroup.Group("/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/register", authHandler.Register)
	}

	var systemGroup = publicGroup.Group("/system")
	{
		systemGroup.GET("/status", statusHandler.Get)
		var patch = systemGroup.Group("/patch")
		{
			patch.Use(authHandler.TokenAuthorise)
			patch.GET("/recent", patchHandler.Recent)
			patch.GET("/list", patchHandler.List)
		}
	}

	return &Router{router}, nil

}

func (r *Router) Serve(addr string) error {
	return r.Run(addr)
}
