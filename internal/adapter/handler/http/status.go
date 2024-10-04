package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/core/port"
)

type StatusHandler struct {
	svc port.SystemStatusService
}

func NewStatusHandler(svc port.SystemStatusService) *StatusHandler {
	return &StatusHandler{svc: svc}
}

func (h *StatusHandler) Get(ctx *gin.Context) {

	if err := h.svc.Status(); err != nil {
		ctx.AbortWithError(503, err)
	}

}
