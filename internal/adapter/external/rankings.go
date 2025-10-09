package external

import (
	"context"
	"github.com/justjack1521/mevium/pkg/genproto/protorank"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/player"
)

type RankingRepository struct {
	svc services.MeviusRankServiceClient
}

func (r RankingRepository) QueryTopRankings(ctx context.Context, code string) ([]player.RankPlayer, error) {

	var md = application.MetadataFromContext(ctx)

	result, err := r.svc.GetTopRankings(mevrpc.NewOutgoingContext(ctx, md.UserID, md.PlayerID), &protorank.GetTopRankRequest{EventName: code})
	if err != nil {
		return nil, err
	}

	var results = make([]player.RankPlayer, len(result.Rankings))

	for index, value := range result.Rankings {
		results[index] = player.RankPlayer{
			Rank:  int(value.Rank),
			Score: int64(value.Score),
			Player: player.Player{
				ID:      uuid.FromStringOrNil(value.Identity.PlayerId),
				Name:    value.Identity.PlayerName,
				Level:   int(value.Identity.PlayerLevel),
				Comment: value.Identity.PlayerComment,
			},
			Loadout: player.Loadout{
				JobCardID:       uuid.FromStringOrNil(value.PrimaryLoadout.JobCard.JobCardId),
				SubJobIndex:     int(value.PrimaryLoadout.JobCard.SubJobIndex),
				CrownLevel:      int(value.PrimaryLoadout.JobCard.CrownLevel),
				WeaponID:        uuid.FromStringOrNil(value.PrimaryLoadout.Weapon.WeaponId),
				SubWeaponUnlock: int(value.PrimaryLoadout.Weapon.SubWeaponUnlock),
			},
		}
	}

	return results, nil

}

func NewRankingRepository(svc services.MeviusRankServiceClient) *RankingRepository {
	return &RankingRepository{svc: svc}
}
