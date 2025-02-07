package game

import uuid "github.com/satori/go.uuid"

type AbilityCard struct {
	SysID             uuid.UUID
	Active            bool
	CardNumber        int
	InShop            bool
	BaseCard          BaseCard
	OverrideAbilityID uuid.UUID
	FusionEXPOverride int
	SaleGoldOverride  int
}

type BaseCard struct {
	SysID           uuid.UUID
	Active          bool
	Name            string
	SkillSeedOne    uuid.UUID
	SkillSeedTwo    uuid.UUID
	SkillSeedSplit  string
	SeedFusionBoost int
	AbilityID       uuid.UUID
	Element         uuid.UUID
	Category        string
	FastLearner     bool
}
