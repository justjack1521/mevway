package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/app/web"
	"mevway/internal/decorator"
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

type WebSocketQuery struct {
	SessionID uuid.UUID
	UserID    uuid.UUID
	PlayerID  uuid.UUID
}

type WebSocketHandler decorator.APIRouterHandler[WebSocketQuery]

type webSocketHandler struct {
	server *web.Server
}

func NewWebSocketHandler(srv *web.Server) WebSocketHandler {
	return webSocketHandler{
		server: srv,
	}
}

func (w webSocketHandler) Handle(ctx *gin.Context, query WebSocketQuery) {

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		httperr.BadRequest(err, "Failed to upgrade connection", ctx)
		return
	}

	client := web.NewClient(ctx, w.server, conn)
	client.UserID = query.UserID
	client.PlayerID = query.PlayerID
	client.SessionID = query.SessionID

	w.server.Register <- client

	go client.Read()
	go client.Write()
	go client.Heartbeat()

}
