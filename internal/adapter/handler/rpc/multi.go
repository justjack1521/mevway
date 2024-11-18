package rpc

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protomulti"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	"mevway/internal/core/domain/socket"
)

type MultiServiceClientRouter struct {
	service services.MeviusMultiServiceClient
	routes  map[protomulti.MultiRequestType]handler
}

func NewMultiServiceClientRouter(service services.MeviusMultiServiceClient) *MultiServiceClientRouter {
	var router = &MultiServiceClientRouter{service: service, routes: make(map[protomulti.MultiRequestType]handler)}
	router.routes[protomulti.MultiRequestType_SESSION_CREATE] = router.SessionCreateRoute
	router.routes[protomulti.MultiRequestType_SESSION_END] = router.SessionEndRoute

	router.routes[protomulti.MultiRequestType_LOBBY_CREATE] = router.LobbyCreateRoute
	router.routes[protomulti.MultiRequestType_LOBBY_CANCEL] = router.LobbyCancelRoute
	router.routes[protomulti.MultiRequestType_LOBBY_READY] = router.LobbyReadyRoute
	router.routes[protomulti.MultiRequestType_LOBBY_START] = router.LobbyStartRoute
	router.routes[protomulti.MultiRequestType_LOBBY_STAMP] = router.LobbyStampRoute
	router.routes[protomulti.MultiRequestType_LOBBY_SEARCH] = router.LobbySearchRoute

	router.routes[protomulti.MultiRequestType_PARTICIPANT_JOIN] = router.ParticipantJoinRoute
	router.routes[protomulti.MultiRequestType_PARTICIPANT_LEAVE] = router.ParticipantLeaveRoute
	router.routes[protomulti.MultiRequestType_PARTICIPANT_READY] = router.ParticipantReadyRoute
	router.routes[protomulti.MultiRequestType_PARTICIPANT_UNREADY] = router.ParticipantUnreadyRoute
	router.routes[protomulti.MultiRequestType_PARTICIPANT_WATCH] = router.ParticipantWatchRoute
	router.routes[protomulti.MultiRequestType_PARTICIPANT_UNWATCH] = router.ParticipantUnwatchRoute

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

func (r *MultiServiceClientRouter) SessionCreateRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewSessionCreateRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SessionCreate(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) SessionEndRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewSessionEndRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SessionEnd(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) LobbySearchRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewLobbySearchRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LobbySearch(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) LobbyReadyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewLobbyReadyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LobbyReady(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) LobbyStartRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewLobbyStartRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LobbyStart(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ParticipantWatchRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewParticipantWatchRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ParticipantWatch(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ParticipantUnwatchRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewParticipantUnwatchRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ParticipantUnwatch(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) LobbyCreateRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewLobbyCreateRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LobbyCreate(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) LobbyCancelRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewLobbyCancelRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LobbyCancel(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ParticipantJoinRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewParticipantJoinRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ParticipantJoin(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ParticipantLeaveRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewParticipantLeaveRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ParticipantLeave(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ParticipantReadyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewParticipantReadyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ParticipantReady(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ParticipantUnreadyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewParticipantUnreadyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ParticipantUnready(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) LobbyStampRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protomulti.NewLobbyStampRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.LobbyStamp(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
