package news

import (
	"github.com/google/uuid"
	"grpc/internal/domain/tags"
	"grpc/internal/shared"
	"time"
)

const (
	INACTIVE = iota
	ACTIVE
	DRAFT
)

type Status int

type New struct {
	Id          uuid.UUID
	Title       string
	Text        string
	Status      Status
	Media       []shared.Media
	Attachments []shared.Media
	CreatedAt   time.Time
	UpdatedAt   time.Time
	AcceptedBy  uuid.UUID
	CreatedBy   uuid.UUID
	Tags        []tags.Tag
}
