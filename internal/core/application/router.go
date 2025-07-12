package application

import (
	"context"
	"fmt"
	"log/slog"
	"mevway/internal/core/domain/socket"
	"mevway/internal/core/port"
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
	logger     *slog.Logger
	services   map[socket.ServiceIdentifier]port.SocketMessageRouter
	repository port.ClientRequestRepository
}

func NewServiceRouter(logger *slog.Logger, repository port.ClientRequestRepository) *ServiceRouter {
	return &ServiceRouter{
		logger:     logger,
		services:   make(map[socket.ServiceIdentifier]port.SocketMessageRouter),
		repository: repository,
	}
}

func (r *ServiceRouter) RegisterSubRouter(key int, router port.SocketMessageRouter) {
	r.services[socket.ServiceIdentifier{ID: key}] = router
}

func (r *ServiceRouter) Route(ctx context.Context, message socket.Message) (response socket.Response, err error) {

	var entry = r.logger.With(
		slog.Group("message_attr",
			slog.String("player", message.PlayerID.String()),
			slog.Int("service", message.Service.ID),
			slog.Int("operation", message.Operation.ID),
			slog.Int("bytes", len(message.Data)),
		),
	)

	entry.InfoContext(ctx, "socket message received")

	defer func() {
		if err != nil {
			entry.With("error", err.Error()).ErrorContext(ctx, "socket message route failed")
		} else {
			entry.InfoContext(ctx, "socket message route success")
		}
	}()

	service, exists := r.services[message.Service]
	if exists == false {
		return nil, errFailedRouteRequest(errServiceNotFound(message.Service))
	}

	response, err = service.Route(ctx, message)
	if err != nil {
		return nil, errFailedRouteRequest(err)
	}

	return response, nil

}
