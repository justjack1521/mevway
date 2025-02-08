package game

import uuid "github.com/satori/go.uuid"

type BaseItem struct {
	SysID          uuid.UUID
	Active         bool
	Name           string
	Maximum        int
	MonthlyMaximum int
}
