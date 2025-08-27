package content

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type GameFeature struct {
	SysID     uuid.UUID
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Metrics   []GameFeatureMetric
}

type GameFeatureMetric struct {
	SysID uuid.UUID
	Title string
	Value int
	Total int
}
