package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/core/port"
	"net"
)

type StatusHandler struct {
	svc port.SystemStatusService
}

func NewStatusHandler(svc port.SystemStatusService) *StatusHandler {
	return &StatusHandler{svc: svc}
}

func (h *StatusHandler) Get(ctx *gin.Context) {

	host, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
	if err != nil {
		ctx.AbortWithError(503, err)
		return
	}

	var addr = net.ParseIP(host)
	if addr == nil {
		ctx.AbortWithStatus(503)
		return
	}

	if err := h.svc.Status(addr); err != nil {
		ctx.AbortWithError(503, err)
		return
	}

}
