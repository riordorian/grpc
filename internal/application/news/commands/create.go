package commands

import (
	"context"
	"github.com/google/uuid"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
	"grpc/internal/shared"
	"grpc/internal/shared/dto"

	//"grpc/internal/shared/dto"
	"grpc/internal/shared/interfaces"
)

type CreateHandler struct {
	Repo        news.RepositoryInterface
	Transactor  storage.TransactorInterface
	FileStorage interfaces.FileStorageProviderInterface
}

type CreateHandlerInterface interface {
	Handle(ctx context.Context, req dto.CreateRequest) error
}

func NewCreateHandler(repo news.RepositoryInterface,
	t storage.TransactorInterface,
	f interfaces.FileStorageProviderInterface) CreateHandlerInterface {
	return CreateHandler{
		Repo:        repo,
		Transactor:  t,
		FileStorage: f,
	}
}

func (c CreateHandler) Handle(ctx context.Context, req dto.CreateRequest) error {
	var media []shared.Media
	for i := range req.Media {
		err := c.FileStorage.UploadMediaFromStream(req.Media[i])
		if err != nil {
			return err
		}
		mediaItem := shared.GetMediaInstanceByPath(req.Media[i].GetFileName())
		media = append(media, mediaItem)
	}

	err := c.Transactor.MakeTransaction(ctx, func(ctx context.Context) error {
		id, _ := uuid.Parse("44266dc6-18d0-46bd-a2b5-238de53db2cb")
		newItem := news.News{
			Title:     req.Title,
			Text:      req.Text,
			Status:    1,
			CreatedBy: id,
			Media:     media,
		}
		_, err := c.Repo.Insert(ctx, newItem)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
