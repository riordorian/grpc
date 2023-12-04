package adapters

import (
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters/storage/postgres"
)

type Services struct {
	Database       *postgres.Db
	NewsRepository news.Repository
}

func GetServices() Services {
	db := postgres.GetDb()
	return Services{
		//TODO: add url with confita
		db,
		postgres.GetNewsRepository(db),
	}
}
