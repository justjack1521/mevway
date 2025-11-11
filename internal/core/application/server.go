package application

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"sync"
	"time"
)

const (
	ReapTickerInterval time.Duration = time.Minute * 1
	ReapDuration       time.Duration = 30 * time.Minute
)

type SocketClient struct {
	client   socket.Client
	notifier port.Client
}

type SocketServer struct {
	clientRegister  map[socket.Client]port.Client
	userRegister    map[uuid.UUID]socket.Client
	sessionRegister map[uuid.UUID]uuid.UUID

	register   chan *SocketClient
	unregister chan socket.Client
	notify     chan socket.Message
	terminate  chan uuid.UUID

	publisher *mevent.Publisher

	mu sync.RWMutex
}

func NewSocketServer(publisher *mevent.Publisher) *SocketServer {
	return &SocketServer{
		publisher:       publisher,
		clientRegister:  make(map[socket.Client]port.Client),
		userRegister:    make(map[uuid.UUID]socket.Client),
		sessionRegister: make(map[uuid.UUID]uuid.UUID),
		register:        make(chan *SocketClient),
		unregister:      make(chan socket.Client),
		terminate:       make(chan uuid.UUID),
		notify:          make(chan socket.Message),
	}
}

func (s *SocketServer) Start() {
	go s.run()
	go s.reap()
}

func (s *SocketServer) reap() {

	var ticker = time.NewTicker(ReapTickerInterval)
	defer ticker.Stop()

	for range ticker.C {

		var inactive = make([]port.Client, 0)
		s.mu.Lock()
		var now = time.Now().UTC()

		for _, value := range s.clientRegister {
			if now.Sub(value.LastMessage()) > ReapDuration {
				inactive = append(inactive, value)
			}
		}

		s.mu.Unlock()

		if len(inactive) > 0 {
			s.publisher.Notify(socket.NewServerReapEvent(len(inactive)))
		}

		for _, client := range inactive {
			client.Terminate(socket.ClosureReasonInactivity)
		}

	}

}

func (s *SocketServer) run() {
	for {
		select {
		case c := <-s.register:
			s.HandleRegister(c)
		case c := <-s.unregister:
			s.HandleUnregister(c)
		case c := <-s.terminate:
			s.HandleTerminate(c)
		case n := <-s.notify:
			s.HandleNotify(n)
		}
	}
}

func (s *SocketServer) HandleTerminate(id uuid.UUID) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if client, ok := s.userRegister[id]; ok {
		if notifier, found := s.clientRegister[client]; found {
			go notifier.Close(socket.ClosureReasonServerStop)
			delete(s.clientRegister, client)
			delete(s.userRegister, id)
		}
	}

	delete(s.sessionRegister, id)

	s.publisher.Notify(socket.NewConnectionTerminateEvent(context.Background(), id))

}

func (s *SocketServer) HandleRegister(c *SocketClient) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userID := c.client.UserID

	//if session, exists := s.sessionRegister[userID]; exists {
	//
	//	if c.client.Session == session {
	//
	//		if client, ok := s.userRegister[userID]; ok {
	//			if notifier, found := s.clientRegister[client]; found {
	//				go notifier.Close(socket.ClosureReasonTakeover)
	//				delete(s.clientRegister, client)
	//			}
	//		}
	//
	//	} else {
	//
	//		s.publisher.Notify(socket.NewSuspiciousConnectionEvent(
	//			context.Background(),
	//			userID,
	//			session,
	//			c.client.Session,
	//		))
	//
	//		go c.notifier.Close(socket.ClosureReasonRejected)
	//		return
	//
	//	}
	//}

	s.clientRegister[c.client] = c.notifier
	s.userRegister[userID] = c.client
	s.sessionRegister[userID] = c.client.Session

	s.publisher.Notify(socket.NewClientConnectedEvent(
		context.Background(),
		c.client.Session,
		c.client.UserID,
		c.client.PlayerID,
		c.client.PatchID,
	))

}

func (s *SocketServer) HandleUnregister(c socket.Client) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if value, ok := s.clientRegister[c]; ok {
		s.publisher.Notify(socket.NewClientDisconnectedEvent(
			context.Background(),
			c.Session,
			c.UserID,
			c.PlayerID,
			value.ClosureReason(),
		))

		delete(s.clientRegister, c)

		if current, exists := s.userRegister[c.UserID]; exists && current.Session == c.Session {
			delete(s.userRegister, c.UserID)
		}

	}

}

func (s *SocketServer) HandleNotify(n socket.Message) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for client := range s.clientRegister {
		if client.PlayerID == n.PlayerID {
			s.clientRegister[client].Notify(n.Data)
			break
		}
	}
}

func (s *SocketServer) Register(client socket.Client, notifier port.Client) error {

	select {
	case s.register <- &SocketClient{client: client, notifier: notifier}:
		return nil
	default:
		return nil
	}
}

func (s *SocketServer) Unregister(client socket.Client) {
	select {
	case s.unregister <- client:
		return
	default:
		return
	}
}

func (s *SocketServer) Terminate(id uuid.UUID) {
	select {
	case s.terminate <- id:
	default:
	}
}

func (s *SocketServer) Notify(ctx context.Context, message socket.Message) {
	select {
	case s.notify <- message:
		return
	default:
		return
	}
}

func (s *SocketServer) IsUserConnected(id uuid.UUID) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_, exists := s.userRegister[id]
	return exists
}
