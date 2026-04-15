package application

import (
	"context"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"time"

	"github.com/justjack1521/mevium/pkg/mevent"
	uuid "github.com/satori/go.uuid"
)

const (
	ReapTickerInterval time.Duration = time.Minute * 1
	ReapDuration       time.Duration = 30 * time.Minute
)

type SocketClient struct {
	client   socket.Client
	notifier port.Client
}

type connection struct {
	client   socket.Client
	notifier port.Client
}

type SocketServer struct {
	users   map[uuid.UUID]connection
	players map[uuid.UUID]connection
	revoked map[uuid.UUID]time.Time

	register   chan *SocketClient
	unregister chan socket.Client
	notify     chan socket.Message
	terminate  chan uuid.UUID

	publisher *mevent.Publisher
}

func NewSocketServer(publisher *mevent.Publisher) *SocketServer {
	return &SocketServer{
		publisher:  publisher,
		users:      make(map[uuid.UUID]connection),
		players:    make(map[uuid.UUID]connection),
		revoked:    make(map[uuid.UUID]time.Time),
		register:   make(chan *SocketClient, 64),
		unregister: make(chan socket.Client, 64),
		terminate:  make(chan uuid.UUID, 64),
		notify:     make(chan socket.Message, 128),
	}
}

func (s *SocketServer) Start() {
	go s.run()
}

func (s *SocketServer) run() {
	ticker := time.NewTicker(ReapTickerInterval)
	defer ticker.Stop()

	for {
		select {
		case c := <-s.register:
			s.handleRegister(c)

		case c := <-s.unregister:
			s.handleUnregister(c)

		case id := <-s.terminate:
			s.handleTerminate(id)

		case n := <-s.notify:
			s.handleNotify(n)

		case <-ticker.C:
			s.handleReap()
		}
	}
}

func (s *SocketServer) addConnection(conn connection) {
	s.users[conn.client.UserID] = conn
	s.players[conn.client.PlayerID] = conn
}

func (s *SocketServer) removeConnection(conn connection) {
	delete(s.users, conn.client.UserID)
	delete(s.players, conn.client.PlayerID)
}

func (s *SocketServer) revoke(conn connection) {
	s.revoked[conn.client.Session] = time.Now().UTC()
}

func (s *SocketServer) handleRegister(c *SocketClient) {
	var user = c.client.UserID
	var session = c.client.Session

	// Reject if session already revoked
	if _, revoked := s.revoked[session]; revoked {
		c.notifier.Close(socket.ClosureReasonRejected)
		return
	}

	if existing, exists := s.users[user]; exists {

		// 🔍 Optional: still log suspicious
		if existing.client.Session != session {
			s.publisher.Notify(socket.NewSuspiciousConnectionEvent(
				context.Background(),
				user,
				existing.client.Session,
				session,
			))
		}

		// Revoke old session
		s.revoke(existing)

		// Remove + close the old connection
		s.removeConnection(existing)
		existing.notifier.Close(socket.ClosureReasonTakeover)
	}

	conn := connection{
		client:   c.client,
		notifier: c.notifier,
	}

	s.addConnection(conn)

	s.publisher.Notify(socket.NewClientConnectedEvent(
		context.Background(),
		session,
		c.client.UserID,
		c.client.PlayerID,
		c.client.PatchID,
	))
}

func (s *SocketServer) handleUnregister(c socket.Client) {
	conn, exists := s.users[c.UserID]
	if !exists {
		return
	}

	if conn.client.Session != c.Session {
		return
	}

	s.removeConnection(conn)

	s.publisher.Notify(socket.NewClientDisconnectedEvent(
		context.Background(),
		c.Session,
		c.UserID,
		c.PlayerID,
		conn.notifier.ClosureReason(),
	))
}

func (s *SocketServer) handleTerminate(id uuid.UUID) {
	conn, exists := s.users[id]
	if !exists {
		return
	}

	s.removeConnection(conn)

	conn.notifier.Close(socket.ClosureReasonServerStop)

	s.publisher.Notify(socket.NewConnectionTerminateEvent(
		context.Background(),
		id,
	))
}

func (s *SocketServer) handleNotify(n socket.Message) {
	conn, exists := s.players[n.PlayerID]
	if !exists {
		return
	}

	conn.notifier.Notify(n.Data)
}

func (s *SocketServer) handleReap() {
	now := time.Now().UTC()

	var inactive []connection

	for _, conn := range s.users {
		if now.Sub(conn.notifier.LastMessage()) > ReapDuration {
			inactive = append(inactive, conn)
		}
	}

	if len(inactive) > 0 {
		s.publisher.Notify(socket.NewServerReapEvent(len(inactive)))
	}

	for _, conn := range inactive {
		s.removeConnection(conn)
		conn.notifier.Terminate(socket.ClosureReasonInactivity)
	}

	for session, t := range s.revoked {
		if now.Sub(t) > ReapDuration {
			delete(s.revoked, session)
		}
	}

}

func (s *SocketServer) Register(client socket.Client, notifier port.Client) error {
	s.register <- &SocketClient{client: client, notifier: notifier}
	return nil
}

func (s *SocketServer) Unregister(client socket.Client) {
	s.unregister <- client
}

func (s *SocketServer) Terminate(id uuid.UUID) {
	s.terminate <- id
}

func (s *SocketServer) Notify(ctx context.Context, message socket.Message) {
	s.notify <- message
}
