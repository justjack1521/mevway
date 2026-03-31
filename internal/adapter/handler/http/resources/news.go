package resources

import (
	"fmt"
	"mevway/internal/core/domain/content"
	"time"

	uuid "github.com/satori/go.uuid"
)

type NewsArticle struct {
	ID          uuid.UUID              `json:"id"`
	Title       string                 `json:"title"`
	PublishedAt time.Time              `json:"published_at"`
	Containers  []NewsArticleContainer `json:"containers"`
}

func NewNewsArticle(article content.NewsArticle) (NewsArticle, error) {

	var containers = make([]NewsArticleContainer, len(article.Containers))

	for index, container := range article.Containers {
		c, err := NewNewsArticleContainer(container)
		if err != nil {
			return NewsArticle{}, err
		}
		containers[index] = c
	}

	var result = NewsArticle{
		ID:          article.ID,
		Title:       article.Title,
		PublishedAt: article.PublishedAt,
		Containers:  containers,
	}

	return result, nil

}

type NewsArticleContainer struct {
	ID        uuid.UUID `json:"id"`
	Nodes     []any     `json:"nodes"`
	SortOrder int       `json:"sort_order"`
	Direction string    `json:"direction"`
	Align     string    `json:"align"`
	Gap       string    `json:"gap"`
}

func NewNewsArticleContainer(container content.NewsContainer) (NewsArticleContainer, error) {

	nodes, err := NodesToResponse(container.Nodes)
	if err != nil {
		return NewsArticleContainer{}, err
	}

	var result = NewsArticleContainer{
		ID:        container.ID,
		Nodes:     nodes,
		SortOrder: container.SortOrder,
		Direction: container.Direction,
		Align:     container.Align,
		Gap:       container.Gap,
	}

	return result, nil

}

type HeadingResponse struct {
	ID        uuid.UUID `json:"id"`
	SortOrder int       `json:"sort_order"`
	Type      string    `json:"type"`
	Text      string    `json:"text"`
	Level     int       `json:"level"`
}

type TextResponse struct {
	ID        uuid.UUID `json:"id"`
	SortOrder int       `json:"sort_order"`
	Type      string    `json:"type"`
	Body      string    `json:"body"`
	Format    string    `json:"format"`
}

type ImageResponse struct {
	ID        uuid.UUID `json:"id"`
	SortOrder int       `json:"sort_order"`
	Type      string    `json:"type"`
	Src       string    `json:"src"`
	Alt       string    `json:"alt"`
	Caption   string    `json:"caption,omitempty"`
	Width     int       `json:"width,omitempty"`
	Height    int       `json:"height,omitempty"`
}

type ButtonResponse struct {
	ID        uuid.UUID `json:"id"`
	SortOrder int       `json:"sort_order"`
	Type      string    `json:"type"`
	Text      string    `json:"text"`
	Href      string    `json:"href"`
	Variant   string    `json:"variant"`
	Target    string    `json:"target"`
}

type VideoResponse struct {
	ID        uuid.UUID `json:"id"`
	SortOrder int       `json:"sort_order"`
	Type      string    `json:"type"`
	Src       string    `json:"src"`
	Poster    string    `json:"poster,omitempty"`
	Autoplay  bool      `json:"autoplay"`
	Loop      bool      `json:"loop"`
	Muted     bool      `json:"muted"`
}

type JobResponse struct {
	ID                 uuid.UUID `json:"id"`
	SortOrder          int       `json:"sort_order"`
	Type               string    `json:"type"`
	JobID              uuid.UUID `json:"job_id"`
	Name               string    `json:"name"`
	JobType            string    `json:"job_type"`
	AbilityName        string    `json:"ability_name"`
	AbilityDescription string    `json:"ability_description"`
}

type AbilityCardResponse struct {
	ID                 uuid.UUID `json:"id"`
	SortOrder          int       `json:"sort_order"`
	Type               string    `json:"type"`
	AbilityCardID      uuid.UUID `json:"ability_card_id"`
	Name               string    `json:"name"`
	CardElement        string    `json:"card_element"`
	AbilityName        string    `json:"ability_name"`
	AbilityDescription string    `json:"ability_description"`
}

// response/convert.go

func NewNode(n content.NewsNode) (any, error) {
	switch n := n.(type) {
	case content.HeadingNode:
		return HeadingResponse{
			ID: n.ID, SortOrder: n.SortOrder, Type: "heading",
			Text: n.Text, Level: n.Level,
		}, nil

	case content.TextNode:
		return TextResponse{
			ID: n.ID, SortOrder: n.SortOrder, Type: "text",
			Body: n.Body, Format: n.Format,
		}, nil

	case content.ImageNode:
		return ImageResponse{
			ID: n.ID, SortOrder: n.SortOrder, Type: "image",
			Src: n.Src, Alt: n.Alt, Caption: n.Caption,
			Width: n.Width, Height: n.Height,
		}, nil

	case content.ButtonNode:
		return ButtonResponse{
			ID: n.ID, SortOrder: n.SortOrder, Type: "button",
			Text: n.Text, Href: n.Href,
			Variant: n.Variant, Target: n.Target,
		}, nil

	case content.VideoNode:
		return VideoResponse{
			ID: n.ID, SortOrder: n.SortOrder, Type: "video",
			Src: n.Src, Poster: n.Poster,
			Autoplay: n.Autoplay, Loop: n.Loop, Muted: n.Muted,
		}, nil
	case content.JobCardNode:
		return JobResponse{
			ID:                 n.ID,
			SortOrder:          n.SortOrder,
			Type:               "job",
			JobID:              n.JobID,
			Name:               n.Name,
			JobType:            n.JobType,
			AbilityName:        n.AbilityName,
			AbilityDescription: n.AbilityDescription,
		}, nil
	case content.AbilityCardNode:
		return AbilityCardResponse{
			ID:                 n.ID,
			SortOrder:          n.SortOrder,
			Type:               "ability_card",
			AbilityCardID:      n.AbilityCardID,
			Name:               n.Name,
			CardElement:        n.CardElement,
			AbilityName:        n.AbilityName,
			AbilityDescription: n.AbilityDescription,
		}, nil
	default:
		return nil, fmt.Errorf("unhandled node type: %T", n)
	}
}

func NodesToResponse(nodes []content.NewsNode) ([]any, error) {
	out := make([]any, 0, len(nodes))
	for _, n := range nodes {
		r, err := NewNode(n)
		if err != nil {
			return nil, err
		}
		out = append(out, r)
	}
	return out, nil
}
