package port

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/game"
)

type GameAdminService interface {
	GrantItem(ctx context.Context, player uuid.UUID, item uuid.UUID, quantity int) error
	CreateSkillPanel(ctx context.Context, job uuid.UUID, page int, panel game.SkillPanel) (bool, error)
	CreateBaseJob(ctx context.Context, job game.BaseJob) (bool, error)
	CreateAugmentMaterials(ctx context.Context, id uuid.UUID, materials []game.AugmentMaterial) error
}

type GameValidationService interface {
	ValidateAbilityCard(ctx context.Context, model game.AbilityCard) error
	ValidateBaseItem(ctx context.Context, model game.BaseItem) error
}
