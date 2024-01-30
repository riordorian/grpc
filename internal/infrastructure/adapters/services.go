package adapters

import (
	"context"
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters/storage/postgres"
)

type Services struct {
	Database       *postgres.Db
	NewsRepository news.RepositoryInterface
}

func GetServices(ctx context.Context) (Services, error) {
	cntx, err := postgres.GetContextDb(ctx)
	if err != nil {
		return Services{}, err
	}

	dbx, err := postgres.GetDb(cntx)

	return Services{
		dbx,
		postgres.GetNewsRepository(dbx),
	}, nil
}
