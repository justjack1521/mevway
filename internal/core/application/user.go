package application

import "mevway/internal/domain/user"

type UserEventTranslator interface {
	Created(event user.CreatedEvent) ([]byte, error)
}
