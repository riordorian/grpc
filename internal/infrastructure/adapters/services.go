package adapters

import (
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters/storage/postgres"
)

type Services struct {
	Database       *postgres.Db
	NewsRepository news.RepositoryInterface
}
