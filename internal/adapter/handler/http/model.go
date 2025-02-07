package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/domain/game"
	"mevway/internal/core/port"
	"net/http"
)

type ModelHandler struct {
	svc port.GameValidationService
}

func NewModelHandler(svc port.GameValidationService) *ModelHandler {
	return &ModelHandler{svc: svc}
}

func (h *ModelHandler) ValidateAbilityCard(ctx *gin.Context) {
	var request = &resources.ValidateAbilityCard{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var model = game.AbilityCard{
		SysID: uuid.FromStringOrNil(request.AbilityCard.SysID),
		BaseCard: game.BaseCard{
			SysID:     uuid.FromStringOrNil(request.AbilityCard.BaseCard.SysID),
			Name:      request.AbilityCard.BaseCard.Name,
			AbilityID: uuid.FromStringOrNil(request.AbilityCard.BaseCard.AbilityID),
		},
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.svc.ValidateAbilityCard(actx, model); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	ctx.JSON(200, gin.H{})

}
