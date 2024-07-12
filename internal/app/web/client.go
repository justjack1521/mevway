package web

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"github.com/newrelic/go-agent/v3/newrelic"
	uuid "github.com/satori/go.uuid"
	"time"
)

const (
	maxMessageSize  int64 = 512
	writeWait             = 10 * time.Second
	pongWait              = 60 * time.Second
	pingPeriod            = (pongWait * 9) / 10
	heartBeatPeriod       = 15 * time.Minute
)

const (
	disconnectionSourceRead  = "read_pump"
	disconnectionSourceWrite = "write_pump"
)

type Client struct {
	context             context.Context
	UserID              uuid.UUID
	PlayerID            uuid.UUID
	ConnectionID        uuid.UUID
	connection          *websocket.Conn
	server              *Server
	heartbeatStarted    bool
	closed              bool
	send                chan []byte
	disconnectionSource string
}

func NewClient(ctx context.Context, server *Server, connection *websocket.Conn) (client *Client) {
	client = &Client{
		context:      ctx,
		ConnectionID: uuid.NewV4(),
		connection:   connection,
		server:       server,
		send:         make(chan []byte),
	}
	client.connection.SetCloseHandler(func(code int, text string) error { client.closed = true; return nil })
	return client
}

func (c *Client) NewClientContext(ctx context.Context, request *protocommon.BaseRequest) *ClientContext {
	return &ClientContext{client: c, context: ctx, request: request}
}

var (
	ErrFailedReadClientRequest = func(err error) error {
		return fmt.Errorf("failed to read client request: %w", err)
	}
	ErrFailedUnmarshalMessage = func(err error) error {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	ErrFailedReadMessage = func(err error) error {
		return fmt.Errorf("failed to read client message: %w", err)
	}
)

func (c *Client) Heartbeat() {

	ticker := time.NewTicker(heartBeatPeriod)

	defer func() {
		ticker.Stop()
	}()

	c.server.publisher.Notify(NewClientHeartbeatEvent(c.context, c.UserID, c.PlayerID, c.connection.RemoteAddr()))

	for {
		select {
		case <-ticker.C:
			if c.closed {
				return
			}
			c.server.publisher.Notify(NewClientHeartbeatEvent(c.context, c.UserID, c.PlayerID, c.connection.RemoteAddr()))
		}
	}

}

func (c *Client) Read() {
	defer func() {
		c.disconnectionSource = disconnectionSourceRead
		c.server.Unregister <- c
		c.connection.Close()
	}()

	c.connection.SetReadLimit(maxMessageSize)
	c.connection.SetReadDeadline(time.Now().Add(pongWait))
	c.connection.SetPongHandler(func(string) error { return c.connection.SetReadDeadline(time.Now().Add(pongWait)) })

	for {

		_, message, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(ErrFailedReadClientRequest(ErrFailedReadMessage(err)))
			}
			return
		}

		var txn = c.server.relic.StartTransaction("socket/read")

		request := &protocommon.BaseRequest{}
		if err := proto.Unmarshal(message, request); err != nil {
			err = ErrFailedReadClientRequest(ErrFailedUnmarshalMessage(err))
			c.server.publisher.Notify(ClientMessageErrorEvent{userID: c.UserID, remoteAddr: c.connection.RemoteAddr(), err: err})
			txn.NoticeError(err)
			txn.End()
			break
		}

		if err := c.server.RouteClientRequest(newrelic.NewContext(c.context, txn), c, request); err != nil {
			err = ErrFailedReadClientRequest(ErrFailedRoutingClientRequest(err))
			c.server.publisher.Notify(ClientMessageErrorEvent{userID: c.UserID, remoteAddr: c.connection.RemoteAddr(), err: err})
			txn.NoticeError(err)
			txn.End()
			break
		}

		txn.End()

	}
}

var (
	ErrFailedWriteClientMessage = func(err error) error {
		return fmt.Errorf("failed to write client message: %w", err)
	}
	ErrFailedCloseClientConnection = func(err error) error {
		return fmt.Errorf("failed to close client connection: %w", err)
	}
)

func (c *Client) Write() {

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		c.disconnectionSource = disconnectionSourceWrite
		ticker.Stop()
		c.connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:

			c.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			var txn = c.server.relic.StartTransaction("socket/write")

			writer, err := c.connection.NextWriter(websocket.BinaryMessage)
			if err != nil {
				txn.NoticeError(err)
				txn.End()
				return
			}

			_, err = writer.Write(message)
			if err != nil {
				txn.NoticeError(err)
				txn.End()
				return
			}

			if err := writer.Close(); err != nil {
				txn.NoticeError(err)
				txn.End()
				return
			}

			txn.End()

		case <-ticker.C:
			c.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
