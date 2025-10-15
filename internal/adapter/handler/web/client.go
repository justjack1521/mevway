package web

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxMessageSize  int64 = 1024 * 10
	writeWait             = 10 * time.Second
	pongWait              = 60 * time.Second
	pingPeriod            = (pongWait * 9) / 10
	heartBeatPeriod       = 15 * time.Minute
)

const (
	disconnectionSourceRead  = "read"
	disconnectionSourceWrite = "write"
)

var (
	errFailedReadMessage = func(err error) error {
		return fmt.Errorf("failed to read client message: %w", err)
	}
	errFailedReadClientRequest = func(err error) error {
		return fmt.Errorf("failed to read client request: %w", err)
	}
	errFailedUnmarshalMessage = func(err error) error {
		return fmt.Errorf("failed to unmarshal client message: %w", err)
	}

	errFailedRouteMessage = func(err error) error {
		return fmt.Errorf("failed to route client message: %w", err)
	}
	errFailedWriteMessage = func(err error) error {
		return fmt.Errorf("failed to write client message: %w", err)
	}
)

type Connection struct {
	*websocket.Conn
	send   chan []byte
	mu     sync.Mutex
	closed atomic.Bool
}

func (c *Connection) Close() {

	if c.closed.Swap(true) {
		return
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	close(c.send)
	c.Conn.Close()

}

func (c *Connection) IsClosed() bool {
	return c.closed.Load()
}

type Client struct {
	client        socket.Client
	router        port.SocketMessageRouter
	server        port.SocketServer
	instrumenter  application.TransactionTracer
	translator    application.MessageTranslator
	connection    *Connection
	lastMessage   time.Time
	closureReason socket.ClosureReason
}

func NewClient(client socket.Client, conn *websocket.Conn, server port.SocketServer, router port.SocketMessageRouter, instrumenter application.TransactionTracer, translator application.MessageTranslator) *Client {
	return &Client{
		client: client,
		connection: &Connection{
			Conn: conn,
			send: make(chan []byte, 16),
		},
		server:       server,
		router:       router,
		instrumenter: instrumenter,
		translator:   translator,
		lastMessage:  time.Now().UTC(),
	}
}

func (c *Client) Terminate(reason socket.ClosureReason) {
	if c.connection.IsClosed() {
		return
	}
	c.closureReason = reason
	c.server.Unregister(c.client)
	c.connection.Close()
}

func (c *Client) Close(reason socket.ClosureReason) {

	if c.connection.IsClosed() {
		return
	}

	defer c.Terminate(reason)

	var deadline = time.Now().UTC().Add(time.Minute)

	if err := c.connection.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), deadline); err != nil {
		return
	}

	if err := c.connection.SetReadDeadline(time.Now().UTC().Add(5 * time.Second)); err != nil {
		return
	}

}

func (c *Client) Notify(data []byte) {
	if c.connection.IsClosed() {
		return
	}
	select {
	case c.connection.send <- data:
	default:
	}
}

func (c *Client) error(message socket.Message, response socket.Response, transaction application.Segment) {
	data, err := response.MarshallBinary()
	if err != nil {
		transaction.NoticeError(err)
		return
	}
	var r = c.translator.Response(message, data)
	send, err := r.MarshallBinary()
	if err != nil {
		transaction.NoticeError(err)
	}
	c.Notify(send)
}

func (c *Client) response(message socket.Message, err error, transaction application.Segment) {
	transaction.NoticeError(err)
	var response = c.translator.Error(message, err)
	send, err := response.MarshallBinary()
	if err != nil {
		transaction.NoticeError(err)
		return
	}
	c.Notify(send)
}

func (c *Client) Read() {

	defer func() {
		c.Terminate(socket.ClosureReasonReadStop)
	}()

	c.connection.SetReadLimit(maxMessageSize)
	c.connection.SetReadDeadline(time.Now().Add(pongWait))
	c.connection.SetPongHandler(c.pong)

	for {

		_, message, err := c.connection.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(errFailedReadClientRequest(errFailedReadMessage(err)))
			}
			return
		}

		c.lastMessage = time.Now().UTC()

		ctx, txn := c.instrumenter.Start(context.Background(), "socket/read")
		txn.AddAttribute("user.id", c.client.UserID.String())
		txn.AddAttribute("player.id", c.client.PlayerID.String())

		request, err := c.translator.Message(c.client, message)
		if err != nil {
			txn.NoticeError(errFailedReadClientRequest(errFailedUnmarshalMessage(err)))
			txn.End()
			continue
		}

		txn.AddAttribute("service.key", request.Service.ID)
		txn.AddAttribute("service.operation", request.Operation.ID)

		result, err := c.router.Route(ctx, request)
		if err != nil {
			c.response(request, err, txn)
		} else {
			c.error(request, result, txn)
		}

		txn.End()

	}

}

func (c *Client) Write() {

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Terminate(socket.ClosureReasonWriteStop)
	}()

	for {

		if c.connection.IsClosed() {
			return
		}

		select {

		case <-ticker.C:
			if c.connection.IsClosed() {
				return
			}
			c.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		case message, ok := <-c.connection.send:

			if ok == false {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if c.connection.IsClosed() {
				return
			}

			c.connection.SetWriteDeadline(time.Now().Add(writeWait))

			_, txn := c.instrumenter.Start(context.Background(), "socket/write")
			txn.AddAttribute("user.id", c.client.UserID.String())
			txn.AddAttribute("player.id", c.client.PlayerID.String())

			writer, err := c.connection.NextWriter(websocket.BinaryMessage)
			if err != nil {
				txn.NoticeError(errFailedWriteMessage(err))
				txn.End()
				return
			}

			_, err = writer.Write(message)
			if err != nil {
				txn.NoticeError(errFailedWriteMessage(err))
				txn.End()
				return
			}

			if err := writer.Close(); err != nil {
				txn.NoticeError(errFailedWriteMessage(err))
				txn.End()
				return
			}

			txn.End()

		}
	}

}

func (c *Client) pong(x string) error {
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}

func (c *Client) ClosureReason() socket.ClosureReason {
	return c.closureReason
}

func (c *Client) LastMessage() time.Time {
	return c.lastMessage
}
