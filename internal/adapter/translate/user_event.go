package translate

import (
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"mevway/internal/core/domain/user"
)

type ProtobufUserEventTranslator struct {
}

func NewProtobufUserEventTranslator() *ProtobufUserEventTranslator {
	return &ProtobufUserEventTranslator{}
}

func (t *ProtobufUserEventTranslator) Created(evt user.CreatedEvent) ([]byte, error) {
	var message = &protocommon.UserCreated{
		UserId:     evt.UserID().String(),
		PlayerId:   evt.PlayerID().String(),
		CustomerId: evt.CustomerID(),
	}
	bytes, err := message.MarshallBinary()
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
