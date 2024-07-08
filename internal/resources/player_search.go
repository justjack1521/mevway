package resources

import uuid "github.com/satori/go.uuid"

type PlayerSearchRequest struct {
	CustomerID string `json:"CustomerID" form:"CustomerID"`
}

type PlayerSearchResponse struct {
	PlayerID      uuid.UUID `json:"PlayerID" form:"PlayerID" binding:"required"`
	PlayerName    string    `json:"PlayerName" form:"PlayerName" binding:"required"`
	PlayerLevel   int       `json:"PlayerLevel" form:"PlayerLevel" binding:"required"`
	PlayerComment string    `json:"PlayerComment" form:"PlayerComment" binding:"required"`
	CompanionID   uuid.UUID `json:"CompanionID" form:"CompanionID" binding:"required"`

	JobCardID       uuid.UUID `json:"JobCardID" form:"JobCardID"`
	SubJobIndex     int       `json:"SubJobIndex" form:"SubJobIndex"`
	WeaponID        uuid.UUID `json:"WeaponID" form:"WeaponID"`
	SubWeaponUnlock int       `json:"SubWeaponUnlock" form:"SubWeaponUnlock"`

	RentalCard *PlayerSearchResponseRentalCard `json:"RentalCard" form:"RentalCard"`
}

type PlayerSearchResponseRentalCard struct {
	AbilityCardID    uuid.UUID `json:"AbilityCardID" binding:"required"`
	AbilityCardLevel int       `json:"AbilityCardLevel"`
	AbilityLevel     int       `json:"AbilityLevel"`
	ExtraSkillUnlock int       `json:"ExtraSkillUnlock"`
	OverBoostLevel   int       `json:"OverBoostLevel"`
}
