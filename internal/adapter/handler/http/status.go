package http

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"mevway/internal/core/port"
	"net"
	"net/http"
	"strings"
)

type StatusHandler struct {
	svc port.SystemStatusService
}

func NewStatusHandler(svc port.SystemStatusService) *StatusHandler {
	return &StatusHandler{svc: svc}
}

func (h *StatusHandler) Get(ctx *gin.Context) {

	var list = make([]net.IP, 0)

	var forward = ctx.GetHeader("X-Forwarded-For")

	if forward == "" {
		ctx.AbortWithError(http.StatusServiceUnavailable, errors.New("ip address not found"))
		return
	}

	var actual = strings.Split(forward, ",")
	if len(actual) == 0 {
		ctx.AbortWithError(http.StatusServiceUnavailable, errors.New("ip address not found"))
		return
	}

	var ip = net.ParseIP(actual[0])
	if ip == nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, errors.New("ip address not found"))
		return
	}
	list = append(list, ip)

	if err := h.svc.Status(list); err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, fmt.Errorf("%s failed to connect to server: %w", ip, err))
		return
	}

}
