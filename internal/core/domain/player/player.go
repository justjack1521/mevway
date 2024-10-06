package player

import uuid "github.com/satori/go.uuid"

type SocialPlayer struct {
	Player
	JobCard
	RentalCard
	CompanionID uuid.UUID
}

type Player struct {
	ID      uuid.UUID
	Name    string
	Level   int
	Comment string
}

type JobCard struct {
	JobCardID   uuid.UUID
	SubJobIndex int
	CrownLevel  int
}

type RentalCard struct {
	CardID           uuid.UUID
	CardLevel        int
	AbilityLevel     int
	ExtraSkillUnlock int
	OverboostLevel   int
}
