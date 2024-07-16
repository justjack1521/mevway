package web

import (
	"github.com/justjack1521/mevium/pkg/genproto/protomulti"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
)

type MultiServiceClientRouter struct {
	service services.MeviusMultiServiceClient
	routes  map[protomulti.MultiRequestType]ServiceClientRouteHandler
}

func NewMultiServiceClientRouter(service services.MeviusMultiServiceClient) *MultiServiceClientRouter {
	var router = &MultiServiceClientRouter{service: service, routes: make(map[protomulti.MultiRequestType]ServiceClientRouteHandler)}
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
	return router
}

func (r *MultiServiceClientRouter) Route(ctx *ClientContext, operation int, bytes []byte) (ClientResponse, error) {
	route, exists := r.routes[protomulti.MultiRequestType(operation)]
	if exists == false {
		return nil, ErrFailedRoutingClientRequest(ErrMalformedRequest)
	}
	return route(ctx, bytes)
}

func (r *MultiServiceClientRouter) GetGameRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewGetGameRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.GetGame(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) CreateSessionRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewCreateSessionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CreateSession(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) EndSessionRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewEndSessionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.EndSession(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) SearchLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewSearchLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SearchLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//func (r *MultiServiceClientRouter) SearchPlayerRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
//	request, err := protomulti.NewGetLobbyPlayerRequest(bytes)
//	if err != nil {
//		return nil, err
//	}
//
//	result, err := r.service.GetPlayer(ctx.context, request)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}

func (r *MultiServiceClientRouter) StartLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewStartLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.StartLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) WatchLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewWatchLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.WatchLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

//func (r *MultiServiceClientRouter) DiscardLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
//	request, err := protomulti.NewDiscardLobbyRequest(bytes)
//	if err != nil {
//		return nil, err
//	}
//
//	result, err := r.service.DiscardLobby(ctx.context, request)
//	if err != nil {
//		return nil, err
//	}
//
//	return result, nil
//}

func (r *MultiServiceClientRouter) CreateLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewCreateLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CreateLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) CancelLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewCancelLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CancelLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) JoinLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewJoinLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.JoinLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) ReadyLobbyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewReadyLobbyRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ReadyLobby(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *MultiServiceClientRouter) SendStampRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protomulti.NewSendStampRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SendStamp(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
