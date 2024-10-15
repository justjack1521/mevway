package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/application"
	"mevway/internal/core/port"
	"net/http"
)

type AdminHandler struct {
	svc port.GameAdminService
}

func NewAdminHandler(svc port.GameAdminService) *AdminHandler {
	return &AdminHandler{svc: svc}
}

func (h *AdminHandler) GrantItem(ctx *gin.Context) {

	var request = &resources.AdminGrantItemRequest{}

	if err := ctx.BindJSON(request); err != nil {
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

	if err := h.svc.GrantItem(application.NewApplicationContext(ctx, md), request.PlayerID, request.ItemID, request.Quantity); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.AdminGrantItemResponse{})

}
