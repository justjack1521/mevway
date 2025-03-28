package http

import (
	"github.com/gin-gonic/gin"
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
