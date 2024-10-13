package player

import uuid "github.com/satori/go.uuid"

type SocialPlayer struct {
	Player
	Loadout
	RentalCard
	CompanionID uuid.UUID
	LastOnline  int64
}

type Player struct {
	ID      uuid.UUID
	Name    string
	Level   int
	Comment string
}

type Loadout struct {
	JobCardID       uuid.UUID
	SubJobIndex     int
	CrownLevel      int
	WeaponID        uuid.UUID
	SubWeaponUnlock int
}

type RentalCard struct {
	CardID           uuid.UUID
	CardLevel        int
	AbilityLevel     int
	ExtraSkillUnlock int
	OverboostLevel   int
}
