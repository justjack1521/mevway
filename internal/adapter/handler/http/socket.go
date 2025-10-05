package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"net/http"
	"time"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   readBufferSize,
	WriteBufferSize:  writeBufferSize,
	HandshakeTimeout: time.Second * 10,
	CheckOrigin:      func(r *http.Request) bool { return true },
}

type SocketHandler struct {
	svc        port.SocketServer
	repository port.ClientConnectionRepository
	factory    application.SocketClientFactory
}

func NewSocketHandler(svc port.SocketServer, clients port.ClientConnectionRepository, factory application.SocketClientFactory) *SocketHandler {
	return &SocketHandler{svc: svc, repository: clients, factory: factory}
}

func (h *SocketHandler) List(ctx *gin.Context) {

	clients, err := h.repository.List(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, resources.NewSocketClientListResponse(clients))

}

func (h *SocketHandler) Join(ctx *gin.Context) {

	if ctx.IsWebsocket() == false {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	session, err := middleware.SessionIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	user, err := middleware.UserIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	player, err := middleware.PlayerIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	patch, err := middleware.PatchIDFromContext(ctx)
	if err != nil {
		fmt.Println(fmt.Sprintf("No patch ID for user %s", player.String()))
		//ctx.AbortWithError(http.StatusInternalServerError, err)
		//return
	}

	var c = socket.NewClient(session, user, player)
	c.PatchID = patch

	client, err := h.factory.Create(ctx, c, conn)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := h.svc.Register(c, client); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	go client.Read()
	go client.Write()

}
