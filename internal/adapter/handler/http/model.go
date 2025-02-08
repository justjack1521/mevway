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

func (h *ModelHandler) ValidateBaseItem(ctx *gin.Context) {

	var request = &resources.ValidateBaseItemRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.ValidateModelResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	var model = game.BaseItem{
		SysID:          uuid.FromStringOrNil(request.BaseItem.SysID),
		Active:         request.BaseItem.Active,
		Name:           request.BaseItem.Name,
		Maximum:        request.BaseItem.Maximum,
		MonthlyMaximum: request.BaseItem.MonthlyMaximum,
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.ValidateModelResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	if err := h.svc.ValidateBaseItem(actx, model); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.ValidateModelResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(200, resources.ValidateModelResponse{Error: false})

}

func (h *ModelHandler) ValidateAbilityCard(ctx *gin.Context) {
	var request = &resources.ValidateAbilityCardRequest{}
	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.ValidateModelResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	var model = game.AbilityCard{
		SysID:             uuid.FromStringOrNil(request.AbilityCard.SysID),
		Active:            request.AbilityCard.Active,
		OverrideAbilityID: uuid.FromStringOrNil(request.AbilityCard.OverrideAbilityID),
		FusionEXPOverride: request.AbilityCard.FusionEXPOverride,
		SaleGoldOverride:  request.AbilityCard.SaleGoldOverride,
		BaseCard: game.BaseCard{
			SysID:           uuid.FromStringOrNil(request.AbilityCard.BaseCard.SysID),
			Active:          request.AbilityCard.BaseCard.Active,
			Name:            request.AbilityCard.BaseCard.Name,
			AbilityID:       uuid.FromStringOrNil(request.AbilityCard.BaseCard.AbilityID),
			SkillSeedOne:    uuid.FromStringOrNil(request.AbilityCard.BaseCard.SkillSeedOne),
			SkillSeedTwo:    uuid.FromStringOrNil(request.AbilityCard.BaseCard.SkillSeedTwo),
			SkillSeedSplit:  request.AbilityCard.BaseCard.SkillSeedSplit,
			SeedFusionBoost: request.AbilityCard.BaseCard.SeedFusionBoost,
			Category:        request.AbilityCard.BaseCard.Category,
			FastLearner:     request.AbilityCard.BaseCard.FastLearner,
		},
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.ValidateModelResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	if err := h.svc.ValidateAbilityCard(actx, model); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.ValidateModelResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(200, resources.ValidateModelResponse{Error: false})

}
