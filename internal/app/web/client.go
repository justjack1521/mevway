package web

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	uuid "github.com/satori/go.uuid"
	"log"
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

	var txn = c.server.relic.StartTransaction("socket/read")

	defer func() {
		c.disconnectionSource = disconnectionSourceRead
		c.server.Unregister <- c
		if err := c.connection.Close(); err != nil {
			txn.NoticeError(err)
		}
		txn.End()
	}()

	c.connection.SetReadLimit(maxMessageSize)
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		txn.NoticeError(err)
	}
	c.connection.SetPongHandler(func(string) error {
		if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			txn.NoticeError(err)
		}
		return nil
	})

	for {
		_, message, err := c.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			log.Printf("error: %v", err)
			break
		}
		fmt.Println(len(message))
		//txn := c.server.relic.StartTransaction("socket/read")
		//
		//_, message, err := c.connection.ReadMessage()
		//
		//if err != nil {
		//	txn.NoticeError(err)
		//	txn.End()
		//	break
		//}
		//
		//request := &protocommon.BaseRequest{}
		//if err := proto.Unmarshal(message, request); err != nil {
		//	txn.NoticeError(err)
		//	txn.End()
		//	break
		//}
		//
		//if err := c.server.RouteClientRequest(newrelic.NewContext(context.Background(), txn), c, request); err != nil {
		//	txn.NoticeError(err)
		//	txn.End()
		//	c.server.publisher.Notify(NewClientMessageErrorEvent(c.ClientID, c.connection.RemoteAddr(), err))
		//	break
		//}
		//
		//txn.End()

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

	var txn = c.server.relic.StartTransaction("socket/write")

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.disconnectionSource = disconnectionSourceWrite
		c.server.Unregister <- c
		if err := c.connection.Close(); err != nil {
			txn.NoticeError(err)
		}
		txn.End()
	}()

	for {
		select {
		case message, ok := <-c.send:

			mtx := c.server.relic.StartTransaction("socket/write/message")

			if err := c.connection.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				mtx.NoticeError(err)
			}
			if !ok {
				if err := c.connection.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					mtx.NoticeError(err)
					mtx.End()
				}
				break
			}

			writer, err := c.connection.NextWriter(websocket.BinaryMessage)
			if err != nil {
				mtx.NoticeError(err)
				mtx.End()
				break
			}

			_, err = writer.Write(message)
			if err != nil {
				mtx.NoticeError(err)
				mtx.End()
				break
			}

			if err := writer.Close(); err != nil {
				mtx.NoticeError(err)
				mtx.End()
				break
			}

		case <-ticker.C:

			ptx := c.server.relic.StartTransaction("socket/ping")

			if err := c.connection.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				ptx.NoticeError(err)
			}
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				ptx.NoticeError(err)
				ptx.End()
				break
			}

		}
	}

}
