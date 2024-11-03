package game

import uuid "github.com/satori/go.uuid"

type AutoAbility struct {
	SysID  uuid.UUID
	Active bool
	Name   string
	Max    int
}
