package news

import (
	"context"
	"github.com/google/uuid"
)

type Request interface{}
type Repository interface {
	GetList(ctx context.Context, request ListRequest) ([]New, error)
	GetById(id uuid.UUID) (New, error)
	Insert(New) (uuid.UUID, error)
	Update(id uuid.UUID, fields New) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type ListRequest struct {
	Sort   int32
	Author int32
	Status Status
	Query  string
	Page   int32
}
