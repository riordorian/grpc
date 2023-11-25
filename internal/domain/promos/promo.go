package promos

import (
	"github.com/google/uuid"
	"grpc/internal/shared"
	"time"
)

const (
	INACTIVE = iota
	ACTIVE
	DRAFT
)

type Status int

type Promo struct {
	Id           uuid.UUID
	Title        string
	Text         string
	ActivePeriod string
	Status       Status
	Media        []shared.Media
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
