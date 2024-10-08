package adapters

import (
	"grpc/internal/domain/repository"
	"grpc/internal/domain/search"
	"grpc/internal/infrastructure/db"
	"grpc/internal/shared/interfaces"
)

type Services struct {
	Database            *db.Db
	NewsRepository      repository.NewsRepositoryInterface
	TagsRepository      repository.TagsRepositoryInterface
	AuthProvider        interfaces.AuthProviderInterface
	FileStorageProvider interfaces.FileStorageProviderInterface
	SearchProvider      search.BaseSearchProviderInterface
}
