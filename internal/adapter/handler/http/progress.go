package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
	"net/http"
)

type ProgressHandler struct {
	svc port.ProgressService
}

func NewProgressHandler(svc port.ProgressService) *ProgressHandler {
	return &ProgressHandler{svc: svc}
}

func (h *ProgressHandler) List(ctx *gin.Context) {

	list, err := h.svc.ListProgress(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(200, resources.NewProgressListResponse(list))

}
