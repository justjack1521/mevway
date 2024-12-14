package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/middleware"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	authHandler *AuthenticationHandler,
	userHandler *UserHandler,
	statusHandler *StatusHandler,
	patchHandler *PatchHandler,
	socketHandler *SocketHandler,
	playerHandler *PlayerHandler,
	adminHandler *AdminHandler,
	contactHandler *ContactHandler,
	middle ...gin.HandlerFunc,
) (*Router, error) {

	var router = gin.New()
	router.HandleMethodNotAllowed = false

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://blankproject.dev"},
		AllowMethods:     []string{"GET", "POST", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.Use(middle...)

	var privateGroup = router.Group("/private", authHandler.AccessTokenAuthorise, middleware.AdminRoleMiddleware())
	{
		var gameGroup = privateGroup.Group("/game")
		{
			gameGroup.POST("/item/grant", adminHandler.GrantItem)
			gameGroup.POST("/job", adminHandler.CreateBaseJob)
			gameGroup.POST("/job/panel", adminHandler.CreateSkillPanel)
		}
		var userGroup = privateGroup.Group("/user")
		{
			userGroup.GET("/list", userHandler.List)
			userGroup.POST("/ban")
			userGroup.POST("/delete", userHandler.Delete)
		}
	}

	var publicGroup = router.Group("/public")
	{
		publicGroup.POST("/contact", contactHandler.CreateContact)
		var socketGroup = publicGroup.Group("/socket", authHandler.AccessTokenAuthorise)
		{
			socketGroup.GET("/join", socketHandler.Join)
			socketGroup.GET("/list", middleware.AdminRoleMiddleware(), socketHandler.List)
		}

		var authGroup = publicGroup.Group("/auth")
		{
			authGroup.POST("/login", authHandler.Login)
			authGroup.POST("/register", userHandler.Register)
		}

		var systemGroup = publicGroup.Group("/system")
		{
			systemGroup.GET("/status", statusHandler.Get)
		}

		var patch = publicGroup.Group("/patch")
		{
			patch.GET("/recent", patchHandler.Recent)
			patch.GET("/list", patchHandler.List)
			patch.GET("/issues", patchHandler.Issues)
		}

		var player = publicGroup.Group("/player")
		{
			player.GET("/me", authHandler.IdentityTokenAuthorise)
			player.GET("/search/:customer_id", authHandler.AccessTokenAuthorise, playerHandler.Search)
		}
	}

	return &Router{router}, nil

}

func (r *Router) Serve(addr string) error {
	return r.Run(addr)
}
