package dto

import (
	"encoding/json"
	"fmt"
	"mevway/internal/core/domain/content"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type ArticleNodeGorm struct {
	SysID       uuid.UUID      `gorm:"primaryKey;column:sys_id"`
	ContainerID uuid.UUID      `gorm:"column:news_item_container"`
	Type        string         `gorm:"column:type"`
	SortOrder   int            `gorm:"column:sort_order"`
	Props       datatypes.JSON `gorm:"type:jsonb;column:props"`
}

func (ArticleNodeGorm) TableName() string {
	return "system.news_item_article_now"
}

type HeadingProps struct {
	Text  string `json:"text"`
	Level int    `json:"level"`
}

type TextProps struct {
	Body   string `json:"body"`
	Format string `json:"format"`
}

type ImageProps struct {
	Src     string `json:"src"`
	Alt     string `json:"alt"`
	Caption string `json:"caption,omitempty"`
	Width   int    `json:"width,omitempty"`
	Height  int    `json:"height,omitempty"`
}

type ButtonProps struct {
	Text    string `json:"text"`
	Href    string `json:"href"`
	Variant string `json:"variant"`
	Target  string `json:"target"`
}

type VideoProps struct {
	Src      string `json:"src"`
	Poster   string `json:"poster,omitempty"`
	Autoplay bool   `json:"autoplay"`
	Loop     bool   `json:"loop"`
	Muted    bool   `json:"muted"`
}

func (n *ArticleNodeGorm) ToEntity() (content.Node, error) {
	switch n.Type {
	case "heading":
		var p HeadingProps
		if err := json.Unmarshal(n.Props, &p); err != nil {
			return nil, fmt.Errorf("heading props: %w", err)
		}
		return content.HeadingNode{
			ID: n.SysID, SortOrder: n.SortOrder,
			Text: p.Text, Level: p.Level,
		}, nil

	case "text":
		var p TextProps
		if err := json.Unmarshal(n.Props, &p); err != nil {
			return nil, fmt.Errorf("text props: %w", err)
		}
		return content.TextNode{
			ID: n.SysID, SortOrder: n.SortOrder,
			Body: p.Body, Format: p.Format,
		}, nil

	case "image":
		var p ImageProps
		if err := json.Unmarshal(n.Props, &p); err != nil {
			return nil, fmt.Errorf("image props: %w", err)
		}
		return content.ImageNode{
			ID: n.SysID, SortOrder: n.SortOrder,
			Src: p.Src, Alt: p.Alt, Caption: p.Caption,
			Width: p.Width, Height: p.Height,
		}, nil

	case "button":
		var p ButtonProps
		if err := json.Unmarshal(n.Props, &p); err != nil {
			return nil, fmt.Errorf("button props: %w", err)
		}
		return content.ButtonNode{
			ID: n.SysID, SortOrder: n.SortOrder,
			Text: p.Text, Href: p.Href,
			Variant: p.Variant, Target: p.Target,
		}, nil

	case "video":
		var p VideoProps
		if err := json.Unmarshal(n.Props, &p); err != nil {
			return nil, fmt.Errorf("video props: %w", err)
		}
		return content.VideoNode{
			ID: n.SysID, SortOrder: n.SortOrder,
			Src: p.Src, Poster: p.Poster,
			Autoplay: p.Autoplay, Loop: p.Loop, Muted: p.Muted,
		}, nil

	default:
		return nil, fmt.Errorf("unknown node type: %s", n.Type)
	}
}
