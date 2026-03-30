package content

import uuid "github.com/satori/go.uuid"

type NewsContainer struct {
	ID        uuid.UUID
	NewsItem  uuid.UUID
	SortOrder int
	Direction string
	Align     string
	Gap       string
	Nodes     []NewsNode
}

type NewsArticle struct {
	ID         uuid.UUID
	Title      string
	Containers []NewsContainer
}

type NewsNode interface {
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

type JobCardNode struct {
	ID                 uuid.UUID
	Name               string
	JobType            string
	AbilityName        string
	AbilityDescription string
}

func (n HeadingNode) nodeType() string { return "heading" }
func (n TextNode) nodeType() string    { return "text" }
func (n ImageNode) nodeType() string   { return "image" }
func (n ButtonNode) nodeType() string  { return "button" }
func (n VideoNode) nodeType() string   { return "video" }
func (n JobCardNode) nodeType() string { return "job" }
