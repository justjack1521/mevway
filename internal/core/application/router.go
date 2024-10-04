package application

import (
	"context"
	"fmt"
	"mevway/internal/core/port"
	"mevway/internal/domain/socket"
)

var (
	errFailedRouteRequest = func(err error) error {
		return fmt.Errorf("failed to route request: %w", err)
	}
	errServiceNotFound = func(key socket.ServiceIdentifier) error {
		return fmt.Errorf("service not found with identifier: %v", key.ID)
	}
)

type ServiceRouter struct {
	services map[socket.ServiceIdentifier]port.SocketMessageRouter
}

func NewServiceRouter() *ServiceRouter {
	return &ServiceRouter{}
}

func (r *ServiceRouter) RegisterSubRouter(key int, router port.SocketMessageRouter) {
	r.services[socket.ServiceIdentifier{ID: key}] = router
}

func (r *ServiceRouter) Route(ctx context.Context, message socket.Message) (socket.Response, error) {

	service, exists := r.services[message.Service]
	if exists == false {
		return nil, errFailedRouteRequest(errServiceNotFound(message.Service))
	}

	response, err := service.Route(ctx, message)
	if err != nil {
		return nil, errFailedRouteRequest(err)
	}

	return response, nil

}
