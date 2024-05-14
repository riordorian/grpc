package news

import (
	"context"
	"github.com/google/uuid"
	"grpc/internal/shared/dto"
)

type Request interface{}
type RepositoryInterface interface {
	GetList(ctx context.Context, request dto.ListRequest) ([]News, error)
	GetById(id uuid.UUID) (News, error)
	Insert(ctx context.Context, fields News) (id uuid.UUID, err error)
	Update(id uuid.UUID, fields News) (bool, error)
	Delete(id uuid.UUID) (bool, error)
}
