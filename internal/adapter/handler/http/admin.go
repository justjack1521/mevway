package http

import (
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/domain/game"
	"mevway/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type AdminHandler struct {
	svc port.GameAdminService
}

func NewAdminHandler(svc port.GameAdminService) *AdminHandler {
	return &AdminHandler{svc: svc}
}

func (h *AdminHandler) CreateBaseJob(ctx *gin.Context) {
	var request = &resources.CreateBaseJobRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := uuid.FromString(request.BaseJobID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	t, err := uuid.FromString(request.TypeID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var job = game.BaseJob{
		SysID:  id,
		Active: request.Active,
		Name:   request.Name,
		Number: request.Number,
		TypeID: t,
	}

	response, err := h.svc.CreateBaseJob(actx, job)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, resources.CreateBaseJobResponse{Created: response})

}

func (h *AdminHandler) CreateAbilityCard(ctx *gin.Context) {

	var request = &resources.CreateAbilityCardRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.CreateAbilityCardResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.CreateAbilityCardResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	var card = game.AbilityCard{
		SysID:      uuid.FromStringOrNil(request.SysID),
		Active:     request.Active,
		CardNumber: request.CardNumber,
		InShop:     request.InShop,
		BaseCard: game.BaseCard{
			SysID: uuid.FromStringOrNil(request.BaseCardID),
		},
		AugmentConfigID:   uuid.FromStringOrNil(request.AugmentConfigID),
		OverrideAbilityID: uuid.FromStringOrNil(request.OverrideAbilityID),
		FusionEXPOverride: request.FusionEXPOverride,
		SaleGoldOverride:  request.SaleGoldOverride,
	}

	if err := h.svc.CreateAbilityCard(actx, card); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resources.CreateAbilityCardResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(200, resources.CreateAbilityCardResponse{})

}

func (h *AdminHandler) CreateBaseCard(ctx *gin.Context) {

	var request = &resources.CreateBaseCardRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.CreateBaseCardResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.CreateBaseCardResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	var card = game.BaseCard{
		SysID:               uuid.FromStringOrNil(request.SysID),
		Active:              request.Active,
		Name:                request.Name,
		SkillSeedOne:        uuid.FromStringOrNil(request.SkillSeedOne),
		SkillSeedTwo:        uuid.FromStringOrNil(request.SkillSeedTwo),
		SkillSeedSplit:      request.SkillSeedSplit,
		SeedFusionBoost:     int(request.SeedFusionBoost),
		EXPFusionMultiplier: request.EXPFusionMultiplier,
		AbilityID:           uuid.FromStringOrNil(request.AbilityID),
		Category:            request.Name,
		FastLearner:         request.FastLearner,
	}

	if err := h.svc.CreateBaseCard(actx, card); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resources.CreateBaseCardResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(200, resources.CreateBaseCardResponse{})

}

func (h *AdminHandler) CreateAugmentMaterials(ctx *gin.Context) {

	var request = &resources.CreateAugmentMaterialsRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.CreateAugmentMaterialsResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, resources.CreateAugmentMaterialsResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	var materials = make([]game.AugmentMaterial, len(request.Materials))
	for index, value := range request.Materials {
		materials[index] = game.AugmentMaterial{
			SysID:    uuid.FromStringOrNil(value.SysID),
			Quantity: value.Quantity,
		}
	}

	if err := h.svc.CreateAugmentMaterials(actx, uuid.FromStringOrNil(request.AbilityCardID), materials); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, resources.CreateAugmentMaterialsResponse{
			Error:        true,
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(200, resources.CreateAugmentMaterialsResponse{})

}

func (h *AdminHandler) CreateSkillPanel(ctx *gin.Context) {

	var request = &resources.CreateSkillPanelRequest{}

	if err := ctx.BindJSON(request); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

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

	actx, err := middleware.ApplicationContextFromMetadata(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := h.svc.GrantItem(actx, request.PlayerID, request.ItemID, request.Quantity); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, resources.AdminGrantItemResponse{})

}
