package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mevway/internal/adapter/handler/http/middleware"
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

	var list = make([]net.IP, 0)

	var ip = middleware.IPFromContext(ctx)
	if ip == nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, errors.New("ip address not found"))
		return
	}
	list = append(list, ip)

	if err := h.svc.Status(ctx, list); err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, fmt.Errorf("%s failed to connect to server: %w", ip, err))
		return
	}

}
