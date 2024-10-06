package application

import (
	"mevway/internal/core/domain/user"
)

type UserEventTranslator interface {
	Created(event user.CreatedEvent) ([]byte, error)
}
