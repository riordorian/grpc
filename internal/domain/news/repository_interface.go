package news

import (
	"context"
	"github.com/google/uuid"
)

type Request interface{}
type RepositoryInterface interface {
	GetList(ctx context.Context, request ListRequest) ([]New, error)
	GetById(id uuid.UUID) (New, error)
	Insert(ctx context.Context, fields New) (id uuid.UUID, err error)
	Update(id uuid.UUID, fields New) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}

type ListRequest struct {
	Sort   string
	Author uuid.UUID
	Status Status
	Query  string
	Page   int32
}
