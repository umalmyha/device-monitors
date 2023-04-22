package model

import (
	"github.com/google/uuid"
	"time"
)

type GetAllDevicesQuery struct {
	Search        *string
	FromLatitude  *float32
	ToLatitude    *float32
	FromLongitude *float32
	ToLongitude   *float32
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
