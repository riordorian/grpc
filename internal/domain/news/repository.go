package news

import "github.com/google/uuid"

type Request interface{}
type Repository interface {
	GetList(request Request) ([]New, error)
	GetById(id uuid.UUID) (New, error)
	Insert(New) (uuid.UUID, error)
	Update(id uuid.UUID, fields New) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}
