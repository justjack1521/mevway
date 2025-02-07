package game

import uuid "github.com/satori/go.uuid"

type AbilityCard struct {
	SysID    uuid.UUID
	BaseCard BaseCard
}

type BaseCard struct {
	SysID     uuid.UUID
	Name      string
	AbilityID uuid.UUID
}
