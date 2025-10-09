package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/player"
)

type ListRankingResponse struct {
	Rankings []PlayerRankingResponse `json:"Rankings"`
}

func NewListRankingResponse(rankings []player.RankPlayer) ListRankingResponse {
	var response = ListRankingResponse{Rankings: make([]PlayerRankingResponse, len(rankings))}
	for index, value := range rankings {
		response.Rankings[index] = NewPlayerRankingResponse(value)
	}
	return response
}

type PlayerRankingResponse struct {
	Rank    int             `json:"Rank"`
	Score   int64           `json:"Score"`
	Player  PlayerResponse  `json:"Player"`
	Loadout LoadoutResponse `json:"Loadout"`
}

func NewPlayerRankingResponse(ranking player.RankPlayer) PlayerRankingResponse {
	return PlayerRankingResponse{
		Rank:  ranking.Rank,
		Score: ranking.Score,
		Player: PlayerResponse{
			PlayerID:      ranking.Player.ID,
			PlayerName:    ranking.Player.Name,
			PlayerLevel:   ranking.Player.Level,
			PlayerComment: ranking.Player.Comment,
		},
		Loadout: LoadoutResponse{
			Job: LoadoutJobResponse{
				JobCardID:   ranking.JobCardID,
				SubJobIndex: ranking.SubJobIndex,
				CrownLevel:  ranking.CrownLevel,
			},
			Weapon: LoadoutWeaponResponse{
				WeaponID:        ranking.WeaponID,
				SubWeaponUnlock: ranking.SubWeaponUnlock,
			},
		},
	}
}

type PlayerResponse struct {
	PlayerID      uuid.UUID `json:"PlayerID" binding:"required"`
	PlayerName    string    `json:"PlayerName" binding:"required"`
	PlayerLevel   int       `json:"PlayerLevel" binding:"required"`
	PlayerComment string    `json:"PlayerComment" binding:"required"`
}

type LoadoutResponse struct {
	Job    LoadoutJobResponse    `json:"Job"`
	Weapon LoadoutWeaponResponse `json:"Weapon"`
}

type LoadoutJobResponse struct {
	JobCardID   uuid.UUID `json:"JobCardID"`
	SubJobIndex int       `json:"SubJobIndex"`
	CrownLevel  int       `json:"CrownLevel"`
}

type LoadoutWeaponResponse struct {
	WeaponID        uuid.UUID `json:"WeaponID"`
	SubWeaponUnlock int       `json:"SubWeaponUnlock"`
}
