package external

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protosocial"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/player"
)

type SocialPlayerRepository struct {
	svc services.MeviusSocialServiceClient
}

func NewSocialPlayerRepository(svc services.MeviusSocialServiceClient) *SocialPlayerRepository {
	return &SocialPlayerRepository{svc: svc}
}

func (r *SocialPlayerRepository) GetByID(ctx context.Context, id uuid.UUID) (player.SocialPlayer, error) {

	var md = application.MetadataFromContext(ctx)

	search, err := r.svc.PlayerSearch(mevrpc.NewOutgoingContext(ctx, md.UserID, md.PlayerID), &protosocial.PlayerSearchRequest{PlayerId: id.String()})
	if err != nil {
		return player.SocialPlayer{}, err
	}

	return player.SocialPlayer{
		Player: player.Player{
			ID:      uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.PlayerId),
			Name:    search.PlayerInfo.PlayerInfo.PlayerName,
			Level:   int(search.PlayerInfo.PlayerInfo.PlayerLevel),
			Comment: search.PlayerInfo.PlayerInfo.PlayerComment,
		},
		CompanionID:   uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.CompanionId),
		LastOnline:    search.PlayerInfo.LastOnline,
		FollowerCount: int(search.PlayerInfo.FollowerCount),
		Loadout: player.Loadout{
			JobCardID:       uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.JobCardId),
			SubJobIndex:     int(search.PlayerInfo.PlayerInfo.SubJobIndex),
			CrownLevel:      int(search.PlayerInfo.PlayerInfo.CrownLevel),
			WeaponID:        uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.WeaponId),
			SubWeaponUnlock: int(search.PlayerInfo.PlayerInfo.SubWeaponUnlock),
		},
		RentalCard: player.RentalCard{
			CardID:           uuid.FromStringOrNil(search.PlayerInfo.PlayerInfo.RentalCard.AbilityCardId),
			CardLevel:        int(search.PlayerInfo.PlayerInfo.RentalCard.AbilityCardLevel),
			AbilityLevel:     int(search.PlayerInfo.PlayerInfo.RentalCard.AbilityLevel),
			ExtraSkillUnlock: int(search.PlayerInfo.PlayerInfo.RentalCard.ExtraSkillUnlock),
			OverboostLevel:   int(search.PlayerInfo.PlayerInfo.RentalCard.OverBoostLevel),
		},
	}, nil

}
