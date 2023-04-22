package model

import (
	"github.com/google/uuid"
	"time"
)

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
