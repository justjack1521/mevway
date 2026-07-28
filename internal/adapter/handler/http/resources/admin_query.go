package resources

import uuid "github.com/satori/go.uuid"

type AdminQueryRegionMapData struct {
	PlayerID uuid.UUID `json:"PlayerID" binding:"required"`
	RegionID uuid.UUID `json:"RegionID" binding:"required"`
}
