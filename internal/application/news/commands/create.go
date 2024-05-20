package commands

import (
	"context"
	"github.com/google/uuid"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
	"grpc/internal/domain/repository"
	"grpc/internal/shared"
	"grpc/internal/shared/dto"

	//"grpc/internal/shared/dto"
	"grpc/internal/shared/interfaces"
)

type CreateHandler struct {
	NewsRepo    repository.NewsRepositoryInterface
	TagsRepo    repository.TagsRepositoryInterface
	Transactor  storage.TransactorInterface
	FileStorage interfaces.FileStorageProviderInterface
}

type CreateHandlerInterface interface {
	Handle(ctx context.Context, req dto.CreateRequest) error
}

func NewCreateHandler(newsRepo repository.NewsRepositoryInterface,
	tagsRepo repository.TagsRepositoryInterface,
	t storage.TransactorInterface,
	f interfaces.FileStorageProviderInterface) CreateHandlerInterface {
	return CreateHandler{
		NewsRepo:    newsRepo,
		TagsRepo:    tagsRepo,
		Transactor:  t,
		FileStorage: f,
	}
}

func (c CreateHandler) Handle(ctx context.Context, req dto.CreateRequest) error {
	var media []shared.Media
	userId, _ := ctx.Value("userId").(uuid.UUID)
	for i := range req.Media {
		err := c.FileStorage.UploadMediaFromStream(req.Media[i])
		if err != nil {
			return err
		}
		mediaItem := shared.GetMediaInstanceByPath(req.Media[i].GetFileName())
		media = append(media, mediaItem)
	}

	err := c.Transactor.MakeTransaction(ctx, func(ctx context.Context) error {
		newItem := news.News{
			Title:     req.Title,
			Text:      req.Text,
			Status:    1,
			CreatedBy: userId,
			Media:     media,
		}
		createdNewsId, err := c.NewsRepo.Insert(ctx, newItem)
		if err != nil {
			return err
		}

		if len(req.Tags) > 0 {
			err = c.TagsRepo.AddNewsTags(ctx, createdNewsId, userId, req.Tags)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
