package tags

import (
	"github.com/google/uuid"
)

type Tag struct {
	Id     uuid.UUID
	Text   string
	UserId uuid.UUID
}
