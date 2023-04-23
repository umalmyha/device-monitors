package model

import (
	"github.com/google/uuid"
	"time"
)

type GetAllDevicesQuery struct {
	Search        *string  `form:"search"`
	FromLatitude  *float32 `form:"fromLatitude"`
	ToLatitude    *float32 `form:"toLatitude"`
	FromLongitude *float32 `form:"fromLongitude"`
	ToLongitude   *float32 `form:"toLongitude"`
}

type Coords struct {
	Latitude  float32
	Longitude float32
}

type Device struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Coords
	UpdatedAt time.Time
	CreatedAt time.Time
}
