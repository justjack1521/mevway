package dto

import (
	"encoding/json"
	"fmt"
	"mevway/internal/core/domain/content"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type ArticleGorm struct {
	SysID uuid.UUID `gorm:"primaryKey;column:sys_id"`
	Title string    `gorm:"column:title"`
}

func (ArticleGorm) TableName() string {
	return "system.news_item"
}

func (n *ArticleGorm) ToEntity() content.NewsArticle {
	return content.NewsArticle{
		ID:    n.SysID,
		Title: n.Title,
	}
}

type ArticleNodeContainer struct {
	SysID     uuid.UUID `gorm:"primaryKey;column:sys_id"`
	NewsItem  uuid.UUID `gorm:"column:news_item"`
	SortOrder int       `gorm:"column:sort_order"`
	Direction string    `gorm:"column:direction"`
	Align     string    `gorm:"column:align"`
	Gap       string    `gorm:"column:gap"`
}

func (ArticleNodeContainer) TableName() string {
	return "system.news_item_article_container"
}

func (n *ArticleNodeContainer) ToEntity() content.NewsContainer {
	return content.NewsContainer{
		ID:        n.SysID,
		NewsItem:  n.NewsItem,
		SortOrder: n.SortOrder,
		Direction: n.Direction,
		Align:     n.Align,
		Gap:       n.Gap,
	}
}

type ArticleNodeGorm struct {
	SysID       uuid.UUID      `gorm:"primaryKey;column:sys_id"`
	ContainerID uuid.UUID      `gorm:"column:news_item_container"`
	Type        string         `gorm:"column:type"`
	SortOrder   int            `gorm:"column:sort_order"`
	Props       datatypes.JSON `gorm:"type:jsonb;column:props"`
}

func (ArticleNodeGorm) TableName() string {
	return "system.news_item_article_node"
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

type JobCardProps struct {
	ID                 uuid.UUID `json:"id"`
	Name               string    `json:"name"`
	JobType            string    `json:"job_type"`
	AbilityName        string    `json:"ability_name"`
	AbilityDescription string    `json:"ability_description"`
}

func (n *ArticleNodeGorm) ToEntity() (content.NewsNode, error) {
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

	case "job":
		var p JobCardProps
		if err := json.Unmarshal(n.Props, &p); err != nil {
			return nil, fmt.Errorf("job card props: %w", err)
		}
		return content.JobCardNode{
			ID:                 p.ID,
			Name:               p.Name,
			JobType:            p.JobType,
			AbilityName:        p.AbilityName,
			AbilityDescription: p.AbilityDescription,
		}, nil

	default:
		return nil, fmt.Errorf("unknown node type: %s", n.Type)
	}
}
