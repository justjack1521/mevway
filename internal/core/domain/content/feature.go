package content

import uuid "github.com/satori/go.uuid"

type GameFeatureRelease struct {
	SysID     uuid.UUID
	Name      string
	BannerURL string
	Items     []GameFeatureReleaseItem
}

type GameFeatureReleaseItem struct {
	Title   string
	Link    string
	Content string
}
