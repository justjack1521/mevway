package web

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
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
	ClientID            uuid.UUID
	ConnectionID        uuid.UUID
	connection          *websocket.Conn
	server              *Server
	heartbeatStarted    bool
	closed              bool
	send                chan []byte
	disconnectionSource string
}

func NewClient(server *Server, connection *websocket.Conn) (client *Client) {
	client = &Client{
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

	for {
		select {
		case <-ticker.C:
			if c.closed {
				return
			}
			c.server.publisher.Notify(ClientHeartbeatEvent{clientID: c.ClientID, remoteAddr: c.connection.RemoteAddr()})
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
			break
		}

		request := &protocommon.BaseRequest{}
		if err := proto.Unmarshal(message, request); err != nil {
			fmt.Println(ErrFailedReadClientRequest(ErrFailedUnmarshalMessage(err)))
			break
		}

		if err := c.server.RouteClientRequest(context.Background(), c, request); err != nil {
			c.server.publisher.Notify(ClientMessageErrorEvent{clientID: c.ClientID, remoteAddr: c.connection.RemoteAddr(), err: err})
			break
		}

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

			writer, err := c.connection.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}

			_, err = writer.Write(message)
			if err != nil {
				return
			}

			if err := writer.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}

}
