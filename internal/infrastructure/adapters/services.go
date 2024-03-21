package adapters

import (
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/db"
	"grpc/internal/shared/interfaces"
)

type Services struct {
	Database            *db.Db
	NewsRepository      news.RepositoryInterface
	AuthProvider        interfaces.AuthProviderInterface
	FileStorageProvider interfaces.FileStorageProviderInterface
}
