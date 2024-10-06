package application

import (
	"context"
	"github.com/justjack1521/mevium/pkg/mevent"
	socket2 "mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"sync"
)

type SocketClient struct {
	client   socket2.Client
	notifier port.Client
}

type SocketServer struct {
	clients map[socket2.Client]port.Client

	register   chan *SocketClient
	unregister chan socket2.Client
	notify     chan socket2.Message

	publisher *mevent.Publisher

	mu sync.RWMutex
}

func NewSocketServer(publisher *mevent.Publisher) *SocketServer {
	return &SocketServer{
		publisher:  publisher,
		clients:    make(map[socket2.Client]port.Client),
		register:   make(chan *SocketClient),
		unregister: make(chan socket2.Client),
		notify:     make(chan socket2.Message),
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

func (s *SocketServer) Register(client socket2.Client, notifier port.Client) {
	select {
	case s.register <- &SocketClient{client: client, notifier: notifier}:
		s.publisher.Notify(socket2.NewClientConnectedEvent(context.Background(), client.Session, client.UserID, client.PlayerID, nil))
		return
	default:
		return
	}
}

func (s *SocketServer) Unregister(client socket2.Client) {
	select {
	case s.unregister <- client:
		s.publisher.Notify(socket2.NewClientDisconnectedEvent(context.Background(), client.Session, client.UserID, client.PlayerID, nil))
		return
	default:
		return
	}
}

func (s *SocketServer) Notify(ctx context.Context, message socket2.Message) {
	select {
	case s.notify <- message:
		return
	default:
		return
	}
}
