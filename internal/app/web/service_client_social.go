package web

import (
	"github.com/justjack1521/mevium/pkg/genproto/protosocial"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
)

type SocialServiceClientRouter struct {
	service services.MeviusSocialServiceClient
	routes  map[protosocial.SocialRequestType]ServiceClientRouteHandler
}

func NewSocialServiceClientRouter(service services.MeviusSocialServiceClient) ServiceClientRouter {
	var router = &SocialServiceClientRouter{service: service, routes: make(map[protosocial.SocialRequestType]ServiceClientRouteHandler)}
	router.routes[protosocial.SocialRequestType_FOLLOW_PLAYER] = router.FollowPlayerRoute
	router.routes[protosocial.SocialRequestType_UNFOLLOW_PLAYER] = router.UnfollowPlayerRoute
	router.routes[protosocial.SocialRequestType_GET_SOCIAL_DATA] = router.GetSocialDataRoute
	return router
}

func (r *SocialServiceClientRouter) Route(ctx *ClientContext, operation int, bytes []byte) (ClientResponse, error) {
	route, exists := r.routes[protosocial.SocialRequestType(operation)]
	if exists == false {
		return nil, ErrFailedRoutingClientRequest(ErrMalformedRequest)
	}
	return route(ctx, bytes)
}

func (r *SocialServiceClientRouter) GetSocialDataRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protosocial.NewFetchPlayerSocialInfo(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FetchPlayerSocialInfo(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *SocialServiceClientRouter) FollowPlayerRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protosocial.NewFollowPlayerRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FollowPlayer(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *SocialServiceClientRouter) UnfollowPlayerRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protosocial.NewUnfollowPlayerRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.UnfollowPlayer(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
