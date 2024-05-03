package web

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protocommon"
)

type ClientResponse interface {
	MarshallBinary() ([]byte, error)
}

type ClientError interface {
	MarshallBinary() ([]byte, error)
	Error() string
}

type ClientContext struct {
	context context.Context
	client  *Client
	request *protocommon.BaseRequest
}

func (c *ClientContext) NewResponse(data []byte) *protocommon.Response {
	return &protocommon.Response{
		Header: &protocommon.ResponseHeader{
			ClientId:     c.client.UserID.String(),
			ConnectionId: c.client.ConnectionID.String(),
			CommandId:    c.request.Header.CommandId,
			Service:      c.request.Service,
			Operation:    c.request.Operation,
		},
		Data: data,
	}
}

func (c *ClientContext) NewError(code int32, slug string) *protocommon.Response {
	return &protocommon.Response{
		Header: &protocommon.ResponseHeader{
			ClientId:     c.client.UserID.String(),
			ConnectionId: c.client.ConnectionID.String(),
			CommandId:    c.request.Header.CommandId,
			Service:      c.request.Service,
			Operation:    c.request.Operation,
			Error:        true,
			ErrorCode:    code,
			ErrorMessage: slug,
		},
		Data: nil,
	}
}
