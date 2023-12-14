package queries

import (
	"context"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
)

type ListHandler struct {
	Repo       news.Repository
	Transactor storage.Transactor
}

type GetListHandler interface {
	Handle(ctx context.Context, req news.ListRequest) ([]news.New, error)
}

func NewGetListHandler(repo news.Repository, t storage.Transactor) ListHandler {
	return ListHandler{
		Repo:       repo,
		Transactor: t,
	}
}

func (l ListHandler) Handle(ctx context.Context, req news.ListRequest) ([]news.New, error) {
	/*err := l.Transactor.MakeTransaction(ctx, func(ctx context.Context) error {
		list, err := l.Repo.GetList(ctx)
		if err != nil {
			return nil, err
		}
		return nil
	})*/
	list, err := l.Repo.GetList(ctx, req)

	if err != nil {
		return nil, err
	}

	return list, nil
}
