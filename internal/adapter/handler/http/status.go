package http

import (
	"github.com/gin-gonic/gin"
	"mevway/internal/core/port"
	"net"
	"net/http"
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
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var addr = net.ParseIP(host)
	if addr == nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := h.svc.Status(addr); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

}
