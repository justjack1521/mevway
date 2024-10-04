package rpc

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protomulti"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	"mevway/internal/domain/socket"
)

type MultiServiceClientRouter struct {
	service services.MeviusMultiServiceClient
	routes  map[protomulti.MultiRequestType]handler
}

func NewMultiServiceClientRouter(service services.MeviusMultiServiceClient) *MultiServiceClientRouter {
	var router = &MultiServiceClientRouter{service: service, routes: make(map[protomulti.MultiRequestType]handler)}
	router.routes[protomulti.MultiRequestType_CREATE_SESSION] = router.CreateSessionRoute
	router.routes[protomulti.MultiRequestType_END_SESSION] = router.EndSessionRoute
	router.routes[protomulti.MultiRequestType_CREATE_LOBBY] = router.CreateLobbyRoute
	router.routes[protomulti.MultiRequestType_CANCEL_LOBBY] = router.CancelLobbyRoute
	router.routes[protomulti.MultiRequestType_SEARCH_LOBBY] = router.SearchLobbyRoute
	router.routes[protomulti.MultiRequestType_WATCH_LOBBY] = router.WatchLobbyRoute
	router.routes[protomulti.MultiRequestType_JOIN_LOBBY] = router.JoinLobbyRoute
	router.routes[protomulti.MultiRequestType_READY_LOBBY] = router.ReadyLobbyRoute
	router.routes[protomulti.MultiRequestType_SEND_STAMP] = router.SendStampRoute
	router.routes[protomulti.MultiRequestType_START_LOBBY] = router.StartLobbyRoute
	router.routes[protomulti.MultiRequestType_GET_GAME] = router.GetGameRoute
	router.routes[protomulti.MultiRequestType_GAME_READY_PLAYER] = router.GameReadyPlayerRoute
	router.routes[protomulti.MultiRequestType_GAME_ENQUEUE_ACTION] = router.GameEnqueueActionRoute
	router.routes[protomulti.MultiRequestType_GAME_DEQUEUE_ACTION] = router.GameDequeueActionRoute
	router.routes[protomulti.MultiRequestType_GAME_LOCK_ACTIONS] = router.GameLockActionRoute
	return router
}

func (r *MultiServiceClientRouter) Route(ctx context.Context, message socket.Message) (socket.Response, error) {
	var out = mevrpc.NewOutgoingContext(ctx, message.UserID, message.PlayerID)
	route, exists := r.routes[protomulti.MultiRequestType(message.Operation.ID)]
	if exists == false {
		return nil, errFailedRoutingClientRequest(errMalformedRequest)
	}
	return route(out, message.Data)
}

func (r *MultiServiceClientRouter) GameReadyPlayerRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewGameReadyPlayerRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ReadyPlayer(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) GameEnqueueActionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewGameEnqueueAbilityRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.EnqueueAction(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) GameDequeueActionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewGameDequeueAbilityRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.DequeueAction(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) GameLockActionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewGameLockActionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LockAction(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) GetGameRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewGetGameRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.GetGame(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) CreateSessionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewCreateSessionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CreateSession(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) EndSessionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewEndSessionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.EndSession(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) SearchLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewSearchLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SearchLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) StartLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewStartLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.StartLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) WatchLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewWatchLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.WatchLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) CreateLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewCreateLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CreateLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) CancelLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewCancelLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CancelLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) JoinLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewJoinLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.JoinLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ReadyLobbyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewReadyLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ReadyLobby(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) SendStampRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewSendStampRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SendStamp(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
