package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/application"
	"mevway/internal/core/port"
	"net/http"
)

type PlayerHandler struct {
	svc port.PlayerSearchService
}

func NewSearchHandler(svc port.PlayerSearchService) *PlayerHandler {
	return &PlayerHandler{svc: svc}
}

func (h *PlayerHandler) Search(ctx *gin.Context) {

	var request = resources.PlayerSearchRequest{
		CustomerID: ctx.Param("customer_id"),
	}

	user, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	player, err := middleware.PlayerIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var md = application.ContextMetadata{
		UserID:   user,
		PlayerID: player,
	}

	result, err := h.svc.Search(application.NewApplicationContext(ctx, md), request.CustomerID)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewPlayerSearchResponse(result))

}

// Social is Search by player ID: the same public profile, for a client that already knows who
// the player is - from a multiplayer game they shared - and never had their customer ID.
func (h *PlayerHandler) Social(ctx *gin.Context) {

	target, err := uuid.FromString(ctx.Param("player_id"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	player, err := middleware.PlayerIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var md = application.ContextMetadata{
		UserID:   user,
		PlayerID: player,
	}

	result, err := h.svc.Get(application.NewApplicationContext(ctx, md), target)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.NewPlayerSearchResponse(result))

}
