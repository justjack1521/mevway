package web

import (
	"github.com/justjack1521/mevium/pkg/genproto/protogame"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
)

type GameServiceClientRouter struct {
	service services.MeviusGameServiceClient
	routes  map[protogame.GameRequestType]ServiceClientRouteHandler
}

func NewGameServiceClientRouter(service services.MeviusGameServiceClient) ServiceClientRouter {
	var router = &GameServiceClientRouter{service: service, routes: make(map[protogame.GameRequestType]ServiceClientRouteHandler)}
	router.routes[protogame.GameRequestType_GET_PROFILE] = router.FetchProfileRoute
	router.routes[protogame.GameRequestType_CREATE_PROFILE] = router.CreatePlayerProfileRoute
	router.routes[protogame.GameRequestType_UPDATE_PROFILE] = router.UpdateProfileRoute
	router.routes[protogame.GameRequestType_FIRST_DAILY_LOGIN] = router.FirstDailyLoginRoute
	router.routes[protogame.GameRequestType_CARD_SALE] = router.CardSaleRoute
	router.routes[protogame.GameRequestType_CARD_FAVOURITE] = router.CardFavouriteRoute
	router.routes[protogame.GameRequestType_SKILL_PANEL] = router.SkillPanelUnlockRoute
	router.routes[protogame.GameRequestType_DECK_EDIT_ALL] = router.DeckEditAllRoute
	router.routes[protogame.GameRequestType_TELEPORT] = router.TeleportRoute
	router.routes[protogame.GameRequestType_PROCESS_REGION_EVENT] = router.ProcessRegionEventRoute
	router.routes[protogame.GameRequestType_BATTLE_REVIVE] = router.BattleReviveRoute
	router.routes[protogame.GameRequestType_BATTLE_COMPLETE] = router.BattleCompleteRoute
	router.routes[protogame.GameRequestType_CONFIRM_DAILY_MISSION] = router.ConfirmDailyRoute
	router.routes[protogame.GameRequestType_CLAIM_LOGIN_CAMPAIGN] = router.ClaimLoginCampaignRoute
	router.routes[protogame.GameRequestType_CLAIM_EVENT_RANKING] = router.ClaimEventRoute
	router.routes[protogame.GameRequestType_CLAIM_MAILBOX] = router.ClaimMailRoute
	router.routes[protogame.GameRequestType_BATTLE_START] = router.BattleStartRoute
	router.routes[protogame.GameRequestType_CARD_TRANSFER] = router.CardTransferRoute
	router.routes[protogame.GameRequestType_EXPAND_CARD_SLOTS] = router.ExpandCardSlotsRoute
	router.routes[protogame.GameRequestType_CARD_FUSION] = router.CardFusionRoute
	router.routes[protogame.GameRequestType_CARD_FUSION_BOOST] = router.CardBoostFusionRoute
	router.routes[protogame.GameRequestType_STAMINA_RESTORE] = router.StaminaRestoreRoute
	router.routes[protogame.GameRequestType_DELETE_ALL_MAILBOX] = router.DeleteMailRoute
	router.routes[protogame.GameRequestType_CLAIM_RENTAL_REWARD] = router.ClaimRentalRewardRoute
	router.routes[protogame.GameRequestType_PURCHASE_ITEM] = router.PurchaseItemRoute
	router.routes[protogame.GameRequestType_PURCHASE_CARD] = router.PurchaseCardRoute
	router.routes[protogame.GameRequestType_ABILITY_SHOP_PURCHASE] = router.AbilityShopRoute
	router.routes[protogame.GameRequestType_CARD_AUGMENT] = router.CardAugmentRoute
	router.routes[protogame.GameRequestType_CLAIM_DAILY_MISSION] = router.ClaimDailyMissionRoute
	router.routes[protogame.GameRequestType_COMPLETE_REGION_MAP] = router.CompleteRegionMapRoute
	return router
}

func (r *GameServiceClientRouter) Route(ctx *ClientContext, operation int, bytes []byte) (ClientResponse, error) {
	route, exists := r.routes[protogame.GameRequestType(operation)]
	if exists == false {
		return nil, ErrFailedRoutingClientRequest(ErrMalformedRequest)
	}
	return route(ctx, bytes)
}

func (r *GameServiceClientRouter) BattleCompleteRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewBattleCompleteRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.BattleComplete(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) BattleReviveRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewBattleReviveRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.BattleRevive(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) BattleStartRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewBattleStartRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.BattleStart(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardTransferRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCardTransferRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardTransfer(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardBoostFusionRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCardBoostFusionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardBoostFusion(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardFavouriteRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCardFavouriteRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardFavourite(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ExpandCardSlotsRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewExpandAbilityCardSlotRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ExpandAbilityCardSlot(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardFusionRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCardFusionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardFusion(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardSaleRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCardSaleRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardSale(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimEventRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewClaimEventRankingRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimEventRanking(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimDailyMissionRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewClaimDailyMissionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimDailyMission(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimLoginCampaignRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewClaimLoginBonusRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimLoginCampaign(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimMailRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewClaimMailBoxItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimMailboxItem(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) DeleteMailRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewDeleteAllMailboxItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.DeleteAllMailboxItem(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ConfirmDailyRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewConfirmDailyMissionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ConfirmDailyMission(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) DeckEditAllRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewDeckEditAllRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.DeckEdit(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) FetchProfileRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewFetchPlayerDataRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FetchPlayerData(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) StaminaRestoreRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewStaminaRestoreRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.RestoreStamina(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) TeleportRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewTeleportRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.Teleport(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ProcessRegionEventRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewProcessRegionEventRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ProcessRegionEvent(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CreatePlayerProfileRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCreateProfileRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CreateProfile(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) UpdateProfileRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewUpdateProfileRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.UpdateProfile(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) FirstDailyLoginRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewFirstDailyLoginRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FirstDailyLogin(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimRentalRewardRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewClaimRentalCardRewardRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimRentalReward(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) PurchaseItemRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewPurchaseItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.PurchaseItem(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) PurchaseCardRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewPurchaseCardRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.PurchaseCard(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) AbilityShopRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewAbilityShopPurchaseRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.AbilityShopPurchase(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardAugmentRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCardAugmentRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardAugment(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) SkillPanelUnlockRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewSkillPanelUnlockRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SkillPanelUnlock(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CompleteRegionMapRoute(ctx *ClientContext, bytes []byte) (ClientResponse, error) {
	request, err := protogame.NewCompleteRegionMapRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CompleteRegion(ctx.context, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
