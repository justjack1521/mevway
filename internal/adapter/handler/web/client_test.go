package web

import (
	"context"
	"errors"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

type fakeSocketServer struct {
	unregistered atomic.Int32
}

func (f *fakeSocketServer) Register(socket.Client, port.Client) error { return nil }
func (f *fakeSocketServer) Unregister(socket.Client)                  { f.unregistered.Add(1) }
func (f *fakeSocketServer) Notify(context.Context, socket.Message)    {}

// newTestClient returns a Client over a real websocket. Its writer is not started, so nothing
// drains the send queue unless a test does it.
func newTestClient(t *testing.T) (*Client, *fakeSocketServer) {
	t.Helper()

	accepted := make(chan *websocket.Conn, 1)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := (&websocket.Upgrader{}).Upgrade(w, r, nil)
		if err != nil {
			t.Errorf("upgrade: %v", err)
			return
		}
		accepted <- conn
	}))
	t.Cleanup(server.Close)

	peer, _, err := websocket.DefaultDialer.Dial("ws"+strings.TrimPrefix(server.URL, "http"), nil)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(func() { peer.Close() })

	fake := &fakeSocketServer{}
	client := NewClient(socket.Client{}, <-accepted, fake, nil, nil, nil)
	t.Cleanup(func() { client.connection.Close() })
	return client, fake
}

func fillQueue(c *Client) {
	for i := 0; i < cap(c.connection.send); i++ {
		c.Notify([]byte("notification"))
	}
}

func TestReplyIsQueuedWhenThereIsRoom(t *testing.T) {
	client, _ := newTestClient(t)

	if err := client.reply([]byte("reply")); err != nil {
		t.Fatal(err)
	}
	if got := string(<-client.connection.send); got != "reply" {
		t.Fatalf("queued %q, want the reply", got)
	}
}

func TestNotifyDropsWhenQueueIsFull(t *testing.T) {
	client, _ := newTestClient(t)
	fillQueue(client)

	done := make(chan struct{})
	go func() {
		client.Notify([]byte("dropped"))
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("Notify blocked on a full queue")
	}
}

func TestReplyWaitsForRoomInsteadOfDropping(t *testing.T) {
	client, _ := newTestClient(t)
	fillQueue(client)

	go func() {
		time.Sleep(100 * time.Millisecond)
		<-client.connection.send // the writer catches up by one
	}()

	if err := client.reply([]byte("reply")); err != nil {
		t.Fatal(err)
	}

	var last []byte
	for i := 0; i < cap(client.connection.send); i++ {
		last = <-client.connection.send
	}
	if string(last) != "reply" {
		t.Fatalf("last queued message is %q, want the reply", last)
	}
	if client.connection.IsClosed() {
		t.Fatal("connection was closed although room came up in time")
	}
}

func TestReplyClosesStalledConnection(t *testing.T) {
	if testing.Short() {
		t.Skip("waits out replyWait")
	}
	client, fake := newTestClient(t)
	fillQueue(client)

	err := client.reply([]byte("reply"))

	if errors.Is(err, errReplyQueueFull) == false {
		t.Fatalf("got %v, want errReplyQueueFull", err)
	}
	if client.connection.IsClosed() == false {
		t.Fatal("stalled connection was left open")
	}
	if fake.unregistered.Load() != 1 {
		t.Fatalf("client unregistered %d times, want once", fake.unregistered.Load())
	}
}

func TestReplyReturnsWhenConnectionClosesWhileWaiting(t *testing.T) {
	client, _ := newTestClient(t)
	fillQueue(client)

	go func() {
		time.Sleep(100 * time.Millisecond)
		client.connection.Close()
	}()

	start := time.Now()
	if err := client.reply([]byte("reply")); err != nil {
		t.Fatal(err)
	}
	if waited := time.Since(start); waited > 2*time.Second {
		t.Fatalf("reply waited %s after the connection closed", waited)
	}
}

// Closing used to close the send channel while other goroutines could still send on it, which
// panics. Hammer both send paths across a close.
func TestSendingAcrossCloseDoesNotPanic(t *testing.T) {
	client, _ := newTestClient(t)

	go func() {
		for {
			select {
			case <-client.connection.send:
			case <-time.After(time.Second):
				return
			}
		}
	}()

	// Senders keep going until after the close, so some are mid-send when it lands.
	var stop atomic.Bool
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for stop.Load() == false {
				client.Notify([]byte("n"))
				_ = client.reply([]byte("r"))
			}
		}()
	}

	time.Sleep(5 * time.Millisecond)
	client.connection.Close()
	time.Sleep(5 * time.Millisecond)
	stop.Store(true)
	wg.Wait()
}
