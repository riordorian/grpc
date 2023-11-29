package shared

import "github.com/google/uuid"

type MediaType int

const (
	IMAGE = iota
	VIDEO
	DOCUMENT
)

type Media struct {
	Id   uuid.UUID
	Name string
	Path string
	Type MediaType
}
