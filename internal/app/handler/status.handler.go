package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	"mevway/internal/decorator"
	"os"
)

var (
	errServerMaintenance = errors.New("service is undergoing maintenance")
)

type ServerStatus struct {
}

type ServerStatusHandler decorator.APIRouterHandler[ServerStatus]

type serverStatusHandler struct {
}

func NewServerStatusHandler() ServerStatusHandler {
	return serverStatusHandler{}
}

func (h serverStatusHandler) Handle(ctx *gin.Context, query ServerStatus) {
	if os.Getenv("MAINTENANCE_MODE") == "true" {
		httperr.InternalError(errServerMaintenance, "server-under-maintenance", ctx)
	}
}
