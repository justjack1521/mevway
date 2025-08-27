package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/content"
	"time"
)

type ProgressListResponse struct {
	Features []Progress `json:"Features"`
}

func NewProgressListResponse(p []content.GameFeature) ProgressListResponse {
	var response = ProgressListResponse{Features: make([]Progress, len(p))}
	for index, value := range p {
		response.Features[index] = NewProgress(value)
	}
	return response
}

type Progress struct {
	SysID     uuid.UUID        `json:"SysID"`
	Title     string           `json:"Title"`
	CreatedAt time.Time        `json:"CreatedAt"`
	UpdatedAt time.Time        `json:"UpdatedAt"`
	Metrics   []ProgressMetric `json:"Metrics"`
}

func NewProgress(p content.GameFeature) Progress {
	var response = Progress{
		SysID:     p.SysID,
		Title:     p.Title,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		Metrics:   make([]ProgressMetric, len(p.Metrics)),
	}
	for index, value := range p.Metrics {
		response.Metrics[index] = NewProgressMetric(value)
	}
	return response
}

type ProgressMetric struct {
	SysID uuid.UUID `json:"SysID"`
	Title string    `json:"Title"`
	Value int       `json:"Value"`
	Total int       `json:"Total"`
}

func NewProgressMetric(m content.GameFeatureMetric) ProgressMetric {
	return ProgressMetric{
		SysID: m.SysID,
		Title: m.Title,
		Value: m.Value,
		Total: m.Total,
	}
}
