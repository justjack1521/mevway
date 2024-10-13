package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/player"
)

type PlayerSearchRequest struct {
	CustomerID string `json:"CustomerID" form:"CustomerID"`
}

type PlayerSearchResponse struct {
	PlayerID        uuid.UUID                      `json:"PlayerID" binding:"required"`
	PlayerName      string                         `json:"PlayerName" binding:"required"`
	PlayerLevel     int                            `json:"PlayerLevel" binding:"required"`
	PlayerComment   string                         `json:"PlayerComment" binding:"required"`
	CompanionID     uuid.UUID                      `json:"CompanionID" binding:"required"`
	LastOnline      int64                          `json:"LastOnline"`
	JobCardID       uuid.UUID                      `json:"JobCardID"`
	SubJobIndex     int                            `json:"SubJobIndex"`
	CrownLevel      int                            `json:"CrownLevel"`
	WeaponID        uuid.UUID                      `json:"WeaponID"`
	SubWeaponUnlock int                            `json:"SubWeaponUnlock"`
	RentalCard      PlayerSearchResponseRentalCard `json:"RentalCard"`
}

func NewPlayerSearchResponse(player player.SocialPlayer) PlayerSearchResponse {
	var response = PlayerSearchResponse{
		PlayerID:        player.ID,
		PlayerName:      player.Name,
		PlayerLevel:     player.Level,
		PlayerComment:   player.Comment,
		CompanionID:     player.CompanionID,
		JobCardID:       player.JobCardID,
		SubJobIndex:     player.SubJobIndex,
		CrownLevel:      player.CrownLevel,
		WeaponID:        player.WeaponID,
		SubWeaponUnlock: player.SubWeaponUnlock,
		LastOnline:      player.LastOnline,
		RentalCard:      NewPlayerSearchResponseRentalCard(player.RentalCard),
	}
	return response
}

type PlayerSearchResponseRentalCard struct {
	AbilityCardID    uuid.UUID `json:"AbilityCardID" binding:"required"`
	CardLevel        int       `json:"CardLevel"`
	AbilityLevel     int       `json:"AbilityLevel"`
	ExtraSkillUnlock int       `json:"ExtraSkillUnlock"`
	OverBoostLevel   int       `json:"OverBoostLevel"`
}

func NewPlayerSearchResponseRentalCard(card player.RentalCard) PlayerSearchResponseRentalCard {
	return PlayerSearchResponseRentalCard{
		AbilityCardID:    card.CardID,
		CardLevel:        card.CardLevel,
		AbilityLevel:     card.AbilityLevel,
		ExtraSkillUnlock: card.ExtraSkillUnlock,
		OverBoostLevel:   card.OverboostLevel,
	}
}
