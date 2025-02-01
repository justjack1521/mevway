package rpc

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protogame"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	"mevway/internal/core/domain/socket"
)

type GameServiceClientRouter struct {
	service services.MeviusGameServiceClient
	routes  map[protogame.GameRequestType]handler
}

func NewGameServiceClientRouter(service services.MeviusGameServiceClient) *GameServiceClientRouter {
	var router = &GameServiceClientRouter{service: service, routes: make(map[protogame.GameRequestType]handler)}
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
	router.routes[protogame.GameRequestType_CLAIM_ALL_MAILBOX] = router.ClaimAllMailRoute
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
	router.routes[protogame.GameRequestType_COMPLETE_REGION_MAP] = router.ClaimRegionRoute
	router.routes[protogame.GameRequestType_SUMMON_ABILITY_CARD] = router.SummonAbilityCardRoute
	router.routes[protogame.GameRequestType_EXECUTE_DIALOGUE] = router.ExecuteDialogueRoute
	router.routes[protogame.GameRequestType_CLAIM_ITEM_DISTILLER] = router.ClaimDistillerRoute
	router.routes[protogame.GameRequestType_CARD_AUTO_SELL] = router.CardAutoSellRoute
	router.routes[protogame.GameRequestType_UNLOCK_REGION_MAP] = router.UnlockRegionMapRoute
	router.routes[protogame.GameRequestType_UNLOCK_REGION_NODE] = router.UnlockRegionNodeRoute

	router.routes[protogame.GameRequestType_CLAIM_DUNGEON] = router.ClaimDungeonRoute

	return router
}

func (r *GameServiceClientRouter) Route(ctx context.Context, message socket.Message) (socket.Response, error) {
	var out = mevrpc.NewOutgoingContext(ctx, message.UserID, message.PlayerID)
	route, exists := r.routes[protogame.GameRequestType(message.Operation.ID)]
	if exists == false {
		return nil, errFailedRoutingClientRequest(errMalformedRequest)
	}
	return route(out, message.Data)
}

func (r *GameServiceClientRouter) ClaimDungeonRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimDungeonRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimDungeon(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) BattleCompleteRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewBattleCompleteRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.BattleComplete(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) BattleReviveRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewBattleReviveRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.BattleRevive(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) BattleStartRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewBattleStartRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.BattleStart(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardTransferRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardTransferRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardTransfer(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardBoostFusionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardBoostFusionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardBoostFusion(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardFavouriteRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardFavouriteRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardFavourite(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ExpandCardSlotsRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewExpandAbilityCardSlotRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ExpandAbilityCardSlot(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardFusionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardFusionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardFusion(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardSaleRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardSaleRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardSale(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimEventRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimEventRankingRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimEventRanking(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimDailyMissionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimDailyMissionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimDailyMission(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimLoginCampaignRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimLoginBonusRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimLoginCampaign(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimMailRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimMailBoxItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimMailboxItem(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimAllMailRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimAllMailBoxItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimAllMailboxItem(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) DeleteMailRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewDeleteAllMailboxItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.DeleteAllMailboxItem(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ConfirmDailyRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewConfirmDailyMissionRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ConfirmDailyMission(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) DeckEditAllRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewDeckEditAllRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.DeckEdit(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) FetchProfileRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewFetchPlayerDataRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FetchPlayerData(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) StaminaRestoreRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewStaminaRestoreRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.RestoreStamina(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) TeleportRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewTeleportRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.Teleport(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ProcessRegionEventRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewProcessRegionEventRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ProcessRegionEvent(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CreatePlayerProfileRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCreateProfileRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CreateProfile(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) UpdateProfileRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewUpdateProfileRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.UpdateProfile(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) FirstDailyLoginRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewFirstDailyLoginRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.FirstDailyLogin(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimRentalRewardRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimRentalCardRewardRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimRentalReward(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) PurchaseItemRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewPurchaseItemRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.PurchaseItem(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) PurchaseCardRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewPurchaseCardRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.PurchaseCard(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) AbilityShopRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewAbilityShopPurchaseRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.AbilityShopPurchase(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardAugmentRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardAugmentRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardAugment(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) SkillPanelUnlockRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewSkillPanelUnlockRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.SkillPanelUnlock(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimRegionRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimRegionMapRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimRegion(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) SummonAbilityCardRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewAbilityCardSummonRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.AbilityCardSummon(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ExecuteDialogueRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewExecuteDialogueRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ExecuteDialogue(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) ClaimDistillerRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewClaimItemDistillerRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.ClaimItemDistiller(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) CardAutoSellRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewCardAutoSellRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.CardAutoSell(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) UnlockRegionMapRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewRegionMapUnlockRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.UnlockRegion(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GameServiceClientRouter) UnlockRegionNodeRoute(ctx context.Context, bytes []byte) (socket.Response, error) {
	request, err := protogame.NewRegionMapNodeUnlockRequest(bytes)
	if err != nil {
		return nil, err
	}

	result, err := r.service.UnlockRegionNode(ctx, request)
	if err != nil {
		return nil, err
	}

	return result, nil
}
