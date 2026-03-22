package http

import (
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentHandler struct {
	svc port.ProgressService
}

func NewFeatureHandler(svc port.ProgressService) *ContentHandler {
	return &ContentHandler{svc: svc}
}

func (h *ContentHandler) ListProgress(ctx *gin.Context) {

	list, err := h.svc.ListProgress(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, resources.NewProgressListResponse(list))

}

func (h *ContentHandler) ListRelease(ctx *gin.Context) {

	list, err := h.svc.ListRelease(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, resources.NewFeatureReleaseListResponse(list))

}
