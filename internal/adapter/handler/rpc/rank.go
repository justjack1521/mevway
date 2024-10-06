package rpc

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protochallenge"
	"github.com/justjack1521/mevium/pkg/genproto/protorank"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	"mevway/internal/core/domain/socket"
)

type RankServiceClientRouter struct {
	service services.MeviusRankServiceClient
	routes  map[int]handler
}

func NewRankServiceClientRouter(service services.MeviusRankServiceClient) *RankServiceClientRouter {
	var router = &RankServiceClientRouter{service: service, routes: make(map[int]handler)}
	router.routes[int(protorank.RankRequestType_GET_RANKING_EVENTS)] = router.GetRankingEventsRoute
	router.routes[int(protochallenge.ChallengeRequestType_GET_SOCIAL_CHALLENGE)] = router.GetSocialChallengeRoute
	router.routes[int(protochallenge.ChallengeRequestType_JOIN_SOCIAL_CHALLENGE)] = router.JoinSocialChallengeRoute
	return router
}

func (r *RankServiceClientRouter) Route(ctx context.Context, message socket.Message) (socket.Response, error) {
	var out = mevrpc.NewOutgoingContext(ctx, message.UserID, message.PlayerID)
	route, exists := r.routes[int(message.Operation.ID)]
	if exists == false {
		return nil, errFailedRoutingClientRequest(errMalformedRequest)
	}
	return route(out, message.Data)
}

func (r *RankServiceClientRouter) GetRankingEventsRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protorank.NewFetchPlayerRankingInfoRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FetchPlayerRankingInfo(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RankServiceClientRouter) GetSocialChallengeRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protochallenge.NewGetPlayerChallengeRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.GetPlayerChallenge(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *RankServiceClientRouter) JoinSocialChallengeRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protochallenge.NewJoinSocialChallengeRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.JoinSocialChallenge(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
