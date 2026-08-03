package port

import (
	"context"
	"mevway/internal/core/domain/game"
	"net"

	"github.com/justjack1521/mevium/pkg/genproto/protogame"
	uuid "github.com/satori/go.uuid"
)

type GameAdminService interface {
	GrantItem(ctx context.Context, player uuid.UUID, item uuid.UUID, quantity int) error
	CreateSkillPanel(ctx context.Context, job uuid.UUID, page int, panel game.SkillPanel) (bool, error)
	CreateBaseJob(ctx context.Context, job game.BaseJob) (bool, error)
	CreateAbility(ctx context.Context, ability game.Ability) error
	CreateAugmentMaterials(ctx context.Context, id uuid.UUID, materials []game.AugmentMaterial) error
	CreateBaseCard(ctx context.Context, card game.BaseCard) error
	CreateAbilityCard(ctx context.Context, ability game.AbilityCard) error
	QueryRegionData(ctx context.Context, player uuid.UUID, id uuid.UUID) (*protogame.ProtoRegionMapInstance, error)
}

type GameValidationService interface {
	ValidateAbilityCard(ctx context.Context, model game.AbilityCard) error
	ValidateBaseItem(ctx context.Context, model game.BaseItem) error
}

type AdministrationRepository interface {
	IPAddressBlacklisted(ctx context.Context, ip net.IP) (bool, error)
}
