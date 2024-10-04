package http

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"mevway/internal/core/application"
	"mevway/internal/core/port"
	"mevway/internal/domain/socket"
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
	svc     port.SocketServer
	factory application.SocketClientFactory
}

func NewSocketHandler(factory application.SocketClientFactory) *SocketHandler {
	return &SocketHandler{factory: factory}
}

func (h *SocketHandler) Join(ctx *gin.Context) {

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return
	}

	session, err := SessionIDFromContext(ctx)
	if err != nil {
		return
	}

	user, err := UserIDFromContext(ctx)
	if err != nil {
		return
	}

	player, err := PlayerIDFromContext(ctx)
	if err != nil {
		return
	}

	var c = socket.NewClient(session, user, player)
	client, err := h.factory.Create(ctx, c, conn)

	h.svc.Register(c, client)

	go client.Read()
	go client.Write()

}
