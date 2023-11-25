package users

import (
	"github.com/google/uuid"
	"grpc/internal/shared"
)

type User struct {
	Id        uuid.UUID
	LastName  string
	FirstName string
	Position  string
	Img       shared.Media
}
