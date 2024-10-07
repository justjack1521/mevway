package translate

import (
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"mevway/internal/core/domain/socket"
)

type ProtobufSocketEventTranslator struct {
}

func NewProtobufSocketEventTranslator() *ProtobufSocketEventTranslator {
	return &ProtobufSocketEventTranslator{}
}

func (t *ProtobufSocketEventTranslator) Connected(event socket.ClientConnectedEvent) ([]byte, error) {
	var message = &protocommon.ClientConnected{
		SessionId: event.SessionID().String(),
	}
	bytes, err := message.MarshallBinary()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (t *ProtobufSocketEventTranslator) Disconnected(event socket.ClientDisconnectedEvent) ([]byte, error) {
	var message = &protocommon.ClientDisconnected{
		SessionId: event.SessionID().String(),
	}
	bytes, err := message.MarshallBinary()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
