package news

import (
	"context"
	"fmt"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
)

type ListHandler struct {
	Repo       news.Repository
	Transactor storage.Transactor
}

func (l ListHandler) List(ctx context.Context) error {
	a, _ := l.Repo.GetList(ctx)
	fmt.Println(a)
	err := l.Transactor.MakeTransaction(ctx, func(ctx context.Context) error {
		a, _ := l.Repo.GetList(ctx)
		b, _ := l.Repo.GetList(ctx)

		fmt.Println(a, b)

		return nil
	})
	if err != nil {
		return err
	}
	return err
}
