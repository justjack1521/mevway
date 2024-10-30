package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/game"
)

type GameAdminService interface {
	GrantItem(ctx context.Context, player uuid.UUID, item uuid.UUID, quantity int) error
	CreateSkillPanel(ctx context.Context, job uuid.UUID, page int, panel game.SkillPanel) (bool, error)
}
