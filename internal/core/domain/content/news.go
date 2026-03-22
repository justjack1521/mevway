package content

import uuid "github.com/satori/go.uuid"

type NewsArticle struct {
	ID      uuid.UUID
	Heading string
}

type Node interface {
	nodeType() string
}

type HeadingNode struct {
	ID        uuid.UUID
	SortOrder int
	Text      string
	Level     int
}

type TextNode struct {
	ID        uuid.UUID
	SortOrder int
	Body      string
	Format    string
}

type ImageNode struct {
	ID        uuid.UUID
	SortOrder int
	Src       string
	Alt       string
	Caption   string
	Width     int
	Height    int
}

type ButtonNode struct {
	ID        uuid.UUID
	SortOrder int
	Text      string
	Href      string
	Variant   string
	Target    string
}

type VideoNode struct {
	ID        uuid.UUID
	SortOrder int
	Src       string
	Poster    string
	Autoplay  bool
	Loop      bool
	Muted     bool
}

func (n HeadingNode) nodeType() string { return "heading" }
func (n TextNode) nodeType() string    { return "text" }
func (n ImageNode) nodeType() string   { return "image" }
func (n ButtonNode) nodeType() string  { return "button" }
func (n VideoNode) nodeType() string   { return "video" }
