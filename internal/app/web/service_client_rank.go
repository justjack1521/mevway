package web

import (
	"github.com/justjack1521/mevium/pkg/genproto/protochallenge"
	"github.com/justjack1521/mevium/pkg/genproto/protorank"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
)

type RankServiceClientRouter struct {
	service services.MeviusRankServiceClient
	routes  map[int]ServiceClientRouteHandler
}

func NewRankServiceClientRouter(service services.MeviusRankServiceClient) ServiceClientRouter {
	var router = &RankServiceClientRouter{service: service, routes: make(map[int]ServiceClientRouteHandler)}
	router.routes[int(protorank.RankRequestType_GET_RANKING_EVENTS)] = router.GetRankingEventsRoute
	router.routes[int(protochallenge.ChallengeRequestType_GET_SOCIAL_CHALLENGE)] = router.GetSocialChallengeRoute
	router.routes[int(protochallenge.ChallengeRequestType_JOIN_SOCIAL_CHALLENGE)] = router.JoinSocialChallengeRoute
	return router
}

func (r *RankServiceClientRouter) Route(ctx *ClientContext, operation int, bytes []byte) (ClientResponse, error) {
	route, exists := r.routes[operation]
	if exists == false {
		return nil, ErrFailedRoutingClientRequest(ErrMalformedRequest)
	}
	return route(ctx, bytes)
}

func (r *RankServiceClientRouter) GetRankingEventsRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protorank.NewFetchPlayerRankingInfoRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FetchPlayerRankingInfo(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RankServiceClientRouter) GetSocialChallengeRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protochallenge.NewGetPlayerChallengeRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.GetPlayerChallenge(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RankServiceClientRouter) JoinSocialChallengeRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protochallenge.NewJoinSocialChallengeRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.JoinSocialChallenge(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
