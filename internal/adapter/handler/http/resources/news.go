package resources

import (
	"fmt"
	"mevway/internal/core/domain/content"

	uuid "github.com/satori/go.uuid"
)

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

// response/convert.go

func NewNode(n content.Node) (any, error) {
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

	default:
		return nil, fmt.Errorf("unhandled node type: %T", n)
	}
}

func NodesToResponse(nodes []content.Node) ([]any, error) {
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
