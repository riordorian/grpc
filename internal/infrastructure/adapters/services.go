package adapters

import (
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters/storage/postgres"
	"grpc/internal/shared/interfaces"
)

type Services struct {
	Database       *postgres.Db
	NewsRepository news.RepositoryInterface
	AuthProvider   interfaces.AuthProviderInterface
}
