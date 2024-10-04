package application

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	"mevway/internal/core/port"
	"mevway/internal/domain/socket"
	"sync"
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
	go s.Run()
}

func (s *SocketServer) Run() {
	for {
		select {
		case c := <-s.register:
			s.mu.Lock()
			s.clients[c.client] = c.notifier
			s.mu.Unlock()
		case c := <-s.unregister:
			s.mu.Lock()
			if _, ok := s.clients[c]; ok {
				delete(s.clients, c)
			}
			s.mu.Unlock()
		case n := <-s.notify:
			s.mu.RLock()
			for client := range s.clients {
				if client.UserID == n.UserID {
					s.clients[client].Notify(n.Data)
				}
			}
			s.mu.RUnlock()
		}
	}
}

func (s *SocketServer) Register(client socket.Client, notifier port.Client) {
	select {
	case s.register <- &SocketClient{client: client, notifier: notifier}:
		s.publisher.Notify(socket.NewClientConnectedEvent(context.Background(), client.Session, client.UserID, client.PlayerID, nil))
		return
	default:
		return
	}
}

func (s *SocketServer) Unregister(client socket.Client) {
	select {
	case s.unregister <- client:
		s.publisher.Notify(socket.NewClientDisconnectedEvent(context.Background(), client.Session, client.UserID, client.PlayerID, nil))
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
