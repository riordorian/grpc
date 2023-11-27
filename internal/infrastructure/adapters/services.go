package adapters

import (
	"grpc/internal/infrastructure/adapters/storage/postgres"
)

type Services struct {
	NewsRepository *postgres.NewsRepository
}

func GetServices() Services {
	return Services{
		//TODO: add url
		NewsRepository: postgres.GetNewsRepository(""),
	}
}
