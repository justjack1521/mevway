package rpc

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protosocial"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	"mevway/internal/domain/socket"
)

type SocialServiceClientRouter struct {
	service services.MeviusSocialServiceClient
	routes  map[protosocial.SocialRequestType]handler
}

func NewSocialServiceClientRouter(service services.MeviusSocialServiceClient) *SocialServiceClientRouter {
	var router = &SocialServiceClientRouter{service: service, routes: make(map[protosocial.SocialRequestType]handler)}
	router.routes[protosocial.SocialRequestType_FOLLOW_PLAYER] = router.FollowPlayerRoute
	router.routes[protosocial.SocialRequestType_UNFOLLOW_PLAYER] = router.UnfollowPlayerRoute
	router.routes[protosocial.SocialRequestType_GET_SOCIAL_DATA] = router.GetSocialDataRoute
	return router
}

func (r *SocialServiceClientRouter) Route(ctx context.Context, message socket.Message) (socket.Response, error) {
	var out = mevrpc.NewOutgoingContext(ctx, message.UserID, message.PlayerID)
	route, exists := r.routes[protosocial.SocialRequestType(message.Operation.ID)]
	if exists == false {
		return nil, errFailedRoutingClientRequest(errMalformedRequest)
	}
	return route(out, message.Data)
}

func (r *SocialServiceClientRouter) GetSocialDataRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protosocial.NewFetchPlayerSocialInfo(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FetchPlayerSocialInfo(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *SocialServiceClientRouter) FollowPlayerRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protosocial.NewFollowPlayerRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FollowPlayer(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *SocialServiceClientRouter) UnfollowPlayerRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protosocial.NewUnfollowPlayerRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.UnfollowPlayer(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
