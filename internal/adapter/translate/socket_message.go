package translate

import (
	"github.com/golang/protobuf/proto"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
	"mevway/internal/core/domain/socket"
)

type ProtobufSocketMessageTranslator struct {
}

func NewProtobufSocketMessageTranslator() *ProtobufSocketMessageTranslator {
	return &ProtobufSocketMessageTranslator{}
}

func (t *ProtobufSocketMessageTranslator) Translate(client socket.Client, message []byte) (socket.Message, error) {
	var request = &protocommon.BaseRequest{}
	if err := proto.Unmarshal(message, request); err != nil {
		return socket.Message{}, err
	}
	return socket.Message{
		UserID: client.UserID,
		Service: socket.ServiceIdentifier{
			ID: int(request.Service),
		},
		Operation: socket.OperationIdentifier{
			ID: int(request.Operation),
		},
		Data: request.Data,
	}, nil
}

func (t *ProtobufSocketMessageTranslator) Notification(data []byte) (socket.Message, error) {
	notification, err := protocommon.NewNotification(data)
	if err != nil {
		return socket.Message{}, err
	}
	return socket.Message{
		Service: socket.ServiceIdentifier{
			ID: int(notification.Service),
		},
		Operation: socket.OperationIdentifier{
			ID: int(notification.Type),
		},
		Data: data,
	}, nil
}

func (t *ProtobufSocketMessageTranslator) Response(message socket.Message, response []byte) (socket.Response, error) {
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

func (t *ProtobufSocketMessageTranslator) Error(message socket.Message, err error) (socket.Response, error) {
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
