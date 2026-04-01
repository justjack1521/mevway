package game

import uuid "github.com/satori/go.uuid"

type AbilityCard struct {
	SysID             uuid.UUID
	Active            bool
	CardNumber        int
	InShop            bool
	BaseCard          BaseCard
	AugmentConfigID   uuid.UUID
	OverrideAbilityID uuid.UUID
	FusionEXPOverride int
	SaleGoldOverride  int
}

type BaseCard struct {
	SysID               uuid.UUID
	Active              bool
	Name                string
	SkillSeedOne        uuid.UUID
	SkillSeedTwo        uuid.UUID
	SkillSeedSplit      string
	SeedFusionBoost     int
	EXPFusionMultiplier float32
	AbilityID           uuid.UUID
	Category            string
	FastLearner         bool
}

type AugmentMaterial struct {
	SysID    uuid.UUID
	Quantity int
}
