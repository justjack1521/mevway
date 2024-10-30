package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/game"
	"mevway/internal/core/port"
	"net/http"
)

type AdminHandler struct {
	svc port.GameAdminService
}

func NewAdminHandler(svc port.GameAdminService) *AdminHandler {
	return &AdminHandler{svc: svc}
}

func (h *AdminHandler) CreateSkillPanel(ctx *gin.Context) {

	var request = &resources.CreateSkillPanelRequest{}

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

	var actx = application.NewApplicationContext(ctx, md)

	var panel = game.SkillPanel{
		DefinitionType: request.Panel.DefinitionType,
		Index:          request.Panel.Index,
		ReferenceID:    uuid.FromStringOrNil(request.Panel.ReferenceID),
		Value:          request.Panel.Value,
		CostItems:      make([]game.CostItem, len(request.Panel.CostItems)),
	}

	for index, value := range request.Panel.CostItems {
		panel.CostItems[index] = game.CostItem{
			ItemID: uuid.FromStringOrNil(value.ItemID),
			Value:  value.Value,
		}
	}

	response, err := h.svc.CreateSkillPanel(actx, uuid.FromStringOrNil(request.BaseJobID), request.PageIndex, panel)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, resources.CreateSkillPanelResponse{Created: response})

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

	var actx = application.NewApplicationContext(ctx, md)

	if err := h.svc.GrantItem(actx, request.PlayerID, request.ItemID, request.Quantity); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.AdminGrantItemResponse{})

}
