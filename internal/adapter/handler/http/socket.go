package http

import (
	"mevway/internal/adapter/handler/http/middleware"
	"mevway/internal/adapter/handler/http/resources"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

// socketPatchApplication is the patch application the game client connects as, matching the
// default /patch/allow answers for.
const socketPatchApplication = "game"

type SocketHandler struct {
	svc        port.SocketServer
	repository port.ClientConnectionRepository
	factory    application.SocketClientFactory
	patches    port.PatchService
}

func NewSocketHandler(svc port.SocketServer, clients port.ClientConnectionRepository, factory application.SocketClientFactory, patches port.PatchService) *SocketHandler {
	return &SocketHandler{svc: svc, repository: clients, factory: factory, patches: patches}
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

	// The patch is checked here, before the upgrade, so a refusal is still an HTTP status the
	// client can read. The boot check alone only runs once per launch, and only if the client
	// chooses to run it - every connection and reconnect comes through here.
	patch, err := middleware.PatchIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUpgradeRequired)
		return
	}

	allowed, err := h.patches.IsPatchAllowed(ctx, socketPatchApplication, patch)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if allowed == false {
		ctx.AbortWithStatus(http.StatusUpgradeRequired)
		return
	}

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
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

	session, err := middleware.SessionIDFromContext(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
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
