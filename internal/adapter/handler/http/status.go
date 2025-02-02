package http

import (
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

	host, _, err := net.SplitHostPort(ctx.Request.RemoteAddr)
	if err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, err)
		return
	}

	var hostAddr = net.ParseIP(host)
	if hostAddr == nil {
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
		return
	}

	list = append(list, hostAddr)

	var forward = ctx.GetHeader("X-Forwarded-For")

	if forward != "" {
		var ranges = strings.Split(forward, ",")
		for _, addr := range ranges {
			var ip = net.ParseIP(addr)
			if ip == nil {
				continue
			}
			list = append(list, ip)
		}
	}

	if err := h.svc.Status(list); err != nil {
		ctx.AbortWithError(http.StatusServiceUnavailable, err)
		return
	}

}
