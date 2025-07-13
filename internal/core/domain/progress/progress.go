package progress

import uuid "github.com/satori/go.uuid"

type GameFeature struct {
	SysID   uuid.UUID
	Title   string
	Metrics []GameFeatureMetric
}

type GameFeatureMetric struct {
	SysID uuid.UUID
	Title string
	Value int
	Total int
}
