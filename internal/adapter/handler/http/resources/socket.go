package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/socket"
)

type SocketClientListResponse struct {
	Clients []SocketClient `json:"Clients"`
}

func NewSocketClientListResponse(clients []socket.Client) SocketClientListResponse {
	var response = SocketClientListResponse{
		Clients: make([]SocketClient, len(clients)),
	}
	for index, value := range clients {
		response.Clients[index] = NewSocketClient(value)
	}
	return response
}

type SocketClient struct {
	SessionID uuid.UUID `json:"SessionID"`
	UserID    uuid.UUID `json:"UserID"`
	PlayerID  uuid.UUID `json:"PlayerID"`
}

func NewSocketClient(client socket.Client) SocketClient {
	return SocketClient{
		SessionID: client.Session,
		UserID:    client.UserID,
		PlayerID:  client.PlayerID,
	}
}
