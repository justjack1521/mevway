package external

import (
	"context"
	"mevway/internal/core/application"
	"mevway/internal/core/domain/player"

	"github.com/justjack1521/mevium/pkg/genproto/protorank"
	services "github.com/justjack1521/mevium/pkg/genproto/service"
	"github.com/justjack1521/mevrpc"
	uuid "github.com/satori/go.uuid"
)

type RankingRepository struct {
	svc services.MeviusRankServiceClient
}

func (r RankingRepository) QueryPlayerRanking(ctx context.Context, code string, p uuid.UUID) (player.RankPlayer, error) {

	var md = application.MetadataFromContext(ctx)

	result, err := r.svc.GetPlayerRankDetails(mevrpc.NewOutgoingContext(ctx, md.PlayerID, md.PlayerID), &protorank.GetPlayerRankDetailsRequest{
		PlayerId:     p.String(),
		Shortcode:    code,
		WithIdentity: true,
		WithLoadout:  true,
	})

	if err != nil {
		return player.RankPlayer{}, err
	}

	return r.convert(result.Details), nil
}

func (r RankingRepository) QueryRankingRange(ctx context.Context, code string, p uuid.UUID, start, stop int) ([]player.RankPlayer, error) {

	var md = application.MetadataFromContext(ctx)

	result, err := r.svc.GetPlayerRankRangeDetails(mevrpc.NewOutgoingContext(ctx, md.PlayerID, md.UserID), &protorank.GetPlayerRankRangeDetailsRequest{
		PlayerId:     p.String(),
		Shortcode:    code,
		Start:        int32(start),
		Stop:         int32(stop),
		WithIdentity: true,
		WithLoadout:  true,
	})
	if err != nil {
		return nil, err
	}

	var results = make([]player.RankPlayer, len(result.Details))

	for index, value := range result.Details {
		results[index] = r.convert(value)
	}

	return results, nil

}

func (r RankingRepository) convert(value *protorank.ProtoPlayerRankSetDetails) player.RankPlayer {
	return player.RankPlayer{
		Rank:      int(value.Rank),
		Primary:   value.PrimaryScore,
		Secondary: value.SecondaryScore,
		Player: player.Player{
			ID:    uuid.FromStringOrNil(value.PlayerId),
			Name:  value.PlayerName,
			Level: int(value.PlayerLevel),
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

func NewRankingRepository(svc services.MeviusRankServiceClient) *RankingRepository {
	return &RankingRepository{svc: svc}
}
