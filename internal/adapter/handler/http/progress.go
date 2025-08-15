package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type FeatureHandler struct {
	svc port.ProgressService
}

func NewFeatureHandler(svc port.ProgressService) *FeatureHandler {
	return &FeatureHandler{svc: svc}
}

func (h *FeatureHandler) ListProgress(ctx *gin.Context) {

	list, err := h.svc.ListProgress(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, resources.NewProgressListResponse(list))

}

func (h *FeatureHandler) ListRelease(ctx *gin.Context) {

	list, err := h.svc.ListRelease(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, resources.NewFeatureReleaseListResponse(list))

}
