package http

import (
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/port"
)

type PatchHandler struct {
	svc port.PatchService
}

func NewPatchHandler(svc port.PatchService) *PatchHandler {
	return &PatchHandler{svc: svc}
}

func (h *PatchHandler) Recent(ctx *gin.Context) {

	current, err := h.svc.Get(ctx, uuid.Nil)
	if err != nil {
		return
	}

	ctx.JSON(200, resources.NewPatchResponse(current))

}

func (h *PatchHandler) List(ctx *gin.Context) {

	list, err := h.svc.GetList(ctx, uuid.Nil, 5)
	if err != nil {
		return
	}

	ctx.JSON(200, resources.NewPatchListResponse(list))

}
