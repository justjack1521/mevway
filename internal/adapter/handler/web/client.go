package web

import (
	"context"
	"fmt"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	maxMessageSize  int64 = 1024 * 10
	writeWait             = 10 * time.Second
	pongWait              = 60 * time.Second
	pingPeriod            = (pongWait * 9) / 10
	heartBeatPeriod       = 15 * time.Minute

	// replyWait is how long a reply may wait for room in the send queue before the connection
	// is given up on. A queue that stays full this long belongs to a stalled connection, which
	// the writer's own writeWait would end shortly anyway.
	replyWait = 5 * time.Second
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
	errReplyQueueFull = fmt.Errorf("reply waited %s for room in a full send queue; closing the connection", replyWait)
)

// Connection is one client's websocket and its outgoing queue.
//
// send is never closed. Notifications and replies are sent on it from other goroutines, and
// closing a channel that someone may still send on panics - here, on the server loop, which
// would take the whole gateway down. Shutdown is signalled on done instead, which the writer
// and any waiting sender select on.
type Connection struct {
	*websocket.Conn
	send   chan []byte
	done   chan struct{}
	closed atomic.Bool
}

func (c *Connection) Close() {

	if c.closed.Swap(true) {
		return
	}

	close(c.done)
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
			done: make(chan struct{}),
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

	_ = c.connection.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(
			websocket.CloseNormalClosure,
			"",
		),
		time.Now().Add(time.Second),
	)

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

// Notify queues a server-pushed notification. One that finds the queue full is dropped: it is
// a live update, and losing one under load is preferable to holding up the server loop that
// delivers everyone's.
func (c *Client) Notify(data []byte) {
	if c.connection.IsClosed() {
		return
	}
	select {
	case c.connection.send <- data:
	default:
	}
}

// reply queues the answer to a client request. Unlike a notification it is never dropped: the
// client is waiting on it, and a lost reply surfaces there as a timeout indistinguishable from a
// slow server, with nothing on either side to say what happened.
//
// When the queue is full it waits up to replyWait for room. If none comes the connection is
// stalled, so it is closed - the client then sees a disconnect and reconnects, rather than
// waiting out its own timeout for an answer that was thrown away.
func (c *Client) reply(data []byte) error {

	if c.connection.IsClosed() {
		return nil
	}

	select {
	case c.connection.send <- data:
		return nil
	default:
	}

	timer := time.NewTimer(replyWait)
	defer timer.Stop()

	select {
	case c.connection.send <- data:
		return nil
	case <-c.connection.done:
		return nil
	case <-timer.C:
		c.Terminate(socket.ClosureReasonWriteStop)
		return errReplyQueueFull
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
	if err := c.reply(send); err != nil {
		transaction.NoticeError(err)
	}
}

func (c *Client) response(message socket.Message, err error, transaction application.Segment) {
	transaction.NoticeError(err)
	var response = c.translator.Error(message, err)
	send, err := response.MarshallBinary()
	if err != nil {
		transaction.NoticeError(err)
		return
	}
	if err := c.reply(send); err != nil {
		transaction.NoticeError(err)
	}
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

		case <-c.connection.done:
			// Terminate has already sent the close frame before closing the connection.
			return

		case message := <-c.connection.send:

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
