package web

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"sync"
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
	closed bool
}

func (c *Connection) Close() {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed == false {
		close(c.send)
		c.Conn.Close()
		c.closed = true
	}

}

type Client struct {
	client       socket.Client
	router       port.SocketMessageRouter
	server       port.SocketServer
	instrumenter application.TransactionInstrumenter
	translator   application.MessageTranslator
	connection   *Connection
}

func NewClient(client socket.Client, conn *websocket.Conn, server port.SocketServer, router port.SocketMessageRouter, instrumenter application.TransactionInstrumenter, translator application.MessageTranslator) *Client {
	return &Client{
		client: client,
		connection: &Connection{
			Conn: conn,
			send: make(chan []byte),
		},
		server:       server,
		router:       router,
		instrumenter: instrumenter,
		translator:   translator,
	}
}

func (c *Client) Close() {
	c.server.Unregister(c.client)
	c.connection.Close()
}

func (c *Client) Notify(data []byte) {
	select {
	case c.connection.send <- data:
		return
	default:
		return
	}
}

func (c *Client) Read() {

	defer func() {
		c.Close()
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

		ctx, txn := c.instrumenter.Start(context.Background(), "socket/read")

		request, err := c.translator.Message(c.client, message)
		if err != nil {
			txn.NoticeError(errFailedReadClientRequest(errFailedUnmarshalMessage(err)))
			txn.End()
			continue
		}

		var response socket.Response

		result, err := c.router.Route(ctx, request)
		if err != nil {

			txn.NoticeError(errFailedReadClientRequest(errFailedRouteMessage(err)))
			txn.End()
			response = c.translator.Error(request, err)

		} else {

			bytes, err := result.MarshallBinary()
			if err != nil {
				txn.NoticeError(errFailedReadClientRequest(errFailedRouteMessage(err)))
				txn.End()
			}
			response = c.translator.Response(request, bytes)

		}

		send, err := response.MarshallBinary()
		if err != nil {
			txn.NoticeError(errFailedReadClientRequest(errFailedRouteMessage(err)))
			txn.End()
			continue
		}

		c.Notify(send)

		txn.End()

	}

}

func (c *Client) Write() {

	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.Close()
	}()

	for {

		select {

		case <-ticker.C:
			c.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		case message, ok := <-c.connection.send:

			c.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if ok == false {
				c.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_, txn := c.instrumenter.Start(context.Background(), "socket/write")

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
