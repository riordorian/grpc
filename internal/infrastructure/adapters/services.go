package adapters

import (
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters/storage/postgres"
)

type Services struct {
	NewsRepository news.Repository
}

func GetServices() Services {
	return Services{
		//TODO: add url with confita
		postgres.GetNewsRepository(),
	}
}
