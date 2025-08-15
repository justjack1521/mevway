package resources

import (
	uuid "github.com/satori/go.uuid"
	"mevway/internal/core/domain/content"
)

type FeatureReleaseListResponse struct {
	Features []FeatureRelease `json:"Features"`
}

func NewFeatureReleaseListResponse(p []content.GameFeatureRelease) FeatureReleaseListResponse {
	var response = FeatureReleaseListResponse{Features: make([]FeatureRelease, len(p))}
	for index, value := range p {
		response.Features[index] = NewFeatureRelease(value)
	}
	return response
}

type FeatureRelease struct {
	SysID     uuid.UUID            `json:"SysID"`
	Name      string               `json:"Name"`
	BannerURL string               `json:"BannerURL"`
	Items     []FeatureReleaseItem `json:"Items"`
}

func NewFeatureRelease(p content.GameFeatureRelease) FeatureRelease {
	var response = FeatureRelease{
		SysID:     p.SysID,
		Name:      p.Name,
		BannerURL: p.BannerURL,
		Items:     make([]FeatureReleaseItem, len(p.Items)),
	}
	for index, value := range p.Items {
		response.Items[index] = NewFeatureContentItem(value)
	}
	return response
}

type FeatureReleaseItem struct {
	Title   string `json:"Title"`
	Link    string `json:"Link"`
	Content string `json:"Content"`
}

func NewFeatureContentItem(p content.GameFeatureReleaseItem) FeatureReleaseItem {
	return FeatureReleaseItem{
		Title:   p.Title,
		Link:    p.Link,
		Content: p.Content,
	}
}
