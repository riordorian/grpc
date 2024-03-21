package commands

import (
	"fmt"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
	//"grpc/internal/shared/dto"
	"grpc/internal/shared/interfaces"
)

type CreateHandler struct {
	Repo        news.RepositoryInterface
	Transactor  storage.TransactorInterface
	FileStorage interfaces.FileStorageProviderInterface
}

type CreateHandlerInterface interface {
	Handle(req string) ([]news.New, error)
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

func (c CreateHandler) Handle(req string) ([]news.New, error) {
	if err := c.FileStorage.Upload("created.png", req); err != nil {
		return nil, err
	}
	/*err := l.Transactor.MakeTransaction(ctx, func(ctx context.Context) error {
		id, _ := uuid.Parse("44266dc6-18d0-46bd-a2b5-238de53db2cb")
		new := news.New{
			Title:     "New 5",
			Text:      "New 5 text",
			Status:    1,
			CreatedBy: id,
		}
		_, err := l.Repo.Insert(ctx, new)
		if err != nil {
			return err
		}

		new2 := news.New{
			Title:     "New 6",
			Text:      "New 6 text",
			Status:    1,
			CreatedBy: id,
		}
		_, err2 := l.Repo.Insert(ctx, new2)
		if err2 != nil {
			return err2
		}
		return nil
	})*/

	fmt.Println(123)
	var list []news.New

	return list, nil
}
