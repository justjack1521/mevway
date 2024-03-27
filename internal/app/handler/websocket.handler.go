package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/justjack1521/mevium/pkg/server/httperr"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/app/web"
	"mevway/internal/decorator"
	"net/http"
)

const (
	readBufferSize  = 1024
	writeBufferSize = 1024
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  readBufferSize,
	WriteBufferSize: writeBufferSize,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type WebSocketQuery struct {
	ClientID string
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

	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if err != nil {
		httperr.BadRequest(err, "Failed to upgrade connection", ctx)
		return
	}

	id, err := uuid.FromString(query.ClientID)
	if err != nil {
		httperr.BadRequest(err, "Failed to upgrade connection", ctx)
		return
	}

	client := web.NewClient(w.server, conn)
	client.ClientID = id
	w.server.Register <- client

	go client.Write()
	go client.Read()
	go client.Heartbeat()

}
