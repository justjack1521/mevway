package translate

import (
	"github.com/golang/protobuf/proto"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	socket2 "mevway/internal/core/domain/socket"
)

type ProtobufSocketMessageTranslator struct {
}

func NewProtobufSocketMessageTranslator() *ProtobufSocketMessageTranslator {
	return &ProtobufSocketMessageTranslator{}
}

func (t *ProtobufSocketMessageTranslator) Translate(client socket2.Client, message []byte) (socket2.Message, error) {
	var request = &protocommon.BaseRequest{}
	if err := proto.Unmarshal(message, request); err != nil {
		return socket2.Message{}, err
	}
	return socket2.Message{
		UserID: client.UserID,
		Service: socket2.ServiceIdentifier{
			ID: int(request.Service),
		},
		Operation: socket2.OperationIdentifier{
			ID: int(request.Operation),
		},
		Data: request.Data,
	}, nil
}

func (t *ProtobufSocketMessageTranslator) Notification(data []byte) (socket2.Message, error) {
	notification, err := protocommon.NewNotification(data)
	if err != nil {
		return socket2.Message{}, err
	}
	return socket2.Message{
		Service: socket2.ServiceIdentifier{
			ID: int(notification.Service),
		},
		Operation: socket2.OperationIdentifier{
			ID: int(notification.Type),
		},
		Data: data,
	}, nil
}

func (t *ProtobufSocketMessageTranslator) Response(message socket2.Message, response []byte) (socket2.Response, error) {
	return &protocommon.Response{
		Header: &protocommon.ResponseHeader{
			ClientId:     message.UserID.String(),
			ConnectionId: message.SessionID.String(),
			CommandId:    message.CommandID.String(),
			Service:      protocommon.ServiceKey(message.Service.ID),
			Operation:    int32(message.Operation.ID),
		},
		Data: response,
	}, nil
}

func (t *ProtobufSocketMessageTranslator) Error(message socket2.Message, err error) (socket2.Response, error) {
	return &protocommon.Response{
		Header: &protocommon.ResponseHeader{
			ClientId:     message.UserID.String(),
			ConnectionId: message.SessionID.String(),
			CommandId:    message.CommandID.String(),
			Service:      protocommon.ServiceKey(message.Service.ID),
			Operation:    int32(message.Operation.ID),
			Error:        true,
			ErrorMessage: err.Error(),
		},
	}, nil
}
