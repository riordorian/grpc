package repository

import (
	"context"
	"github.com/google/uuid"
)

type TagsRepositoryInterface interface {
	BaseRepositoryInterface
	AddNewsTags(ctx context.Context, newsId uuid.UUID, userId uuid.UUID, tags []string) error
}
