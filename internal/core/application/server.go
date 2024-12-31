package application

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"sync"
	"time"
)

type SocketClient struct {
	client   socket.Client
	notifier port.Client
}

type SocketServer struct {
	clients map[socket.Client]port.Client

	register   chan *SocketClient
	unregister chan socket.Client
	notify     chan socket.Message

	publisher *mevent.Publisher

	mu sync.RWMutex
}

func NewSocketServer(publisher *mevent.Publisher) *SocketServer {
	return &SocketServer{
		publisher:  publisher,
		clients:    make(map[socket.Client]port.Client),
		register:   make(chan *SocketClient),
		unregister: make(chan socket.Client),
		notify:     make(chan socket.Message),
	}
}

func (s *SocketServer) Start() {
	go s.run()
	go s.reap()
}

func (s *SocketServer) reap() {
	ticker := time.NewTicker(time.Minute * 1)
	defer ticker.Stop()

	for range ticker.C {

		var inactive = make([]port.Client, 0)
		s.mu.Lock()
		var now = time.Now().UTC()
		for _, value := range s.clients {
			if now.Sub(value.LastMessage()) > 15*time.Minute {
				inactive = append(inactive, value)
			}
		}
		s.mu.Unlock()

		s.publisher.Notify(socket.NewServerReapEvent(len(inactive)))

		for _, client := range inactive {
			client.Close(socket.ClosureReasonServerStop)
		}

	}

}

func (s *SocketServer) run() {
	for {
		select {
		case c := <-s.register:
			s.mu.Lock()
			s.clients[c.client] = c.notifier
			s.publisher.Notify(socket.NewClientConnectedEvent(context.Background(), c.client.Session, c.client.UserID, c.client.PlayerID))
			s.mu.Unlock()
		case c := <-s.unregister:
			s.mu.Lock()
			if value, ok := s.clients[c]; ok {
				s.publisher.Notify(socket.NewClientDisconnectedEvent(context.Background(), c.Session, c.UserID, c.PlayerID, value.ClosureReason()))
				delete(s.clients, c)
			}
			s.mu.Unlock()
		case n := <-s.notify:
			s.mu.RLock()
			for client := range s.clients {
				if client.PlayerID == n.PlayerID {
					s.clients[client].Notify(n.Data)
					break
				}
			}
			s.mu.RUnlock()
		}
	}
}

func (s *SocketServer) Register(client socket.Client, notifier port.Client) error {

	//s.mu.Lock()
	//
	//var connected = false
	//for key := range s.clients {
	//	if key.UserID == client.UserID {
	//		connected = true
	//		break
	//	}
	//}
	//s.mu.Unlock()

	//if connected {
	//	return errors.New("only one session is allowed per user")
	//}

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

func (s *SocketServer) Notify(ctx context.Context, message socket.Message) {
	select {
	case s.notify <- message:
		return
	default:
		return
	}
}
