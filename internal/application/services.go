package application

import (
	appnews "grpc/internal/application/news/queries"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
	"grpc/internal/shared/interfaces"
)

type Services struct {
	//Handler    appnews.ListHandler
	Repository   news.RepositoryInterface
	Transactor   storage.TransactorInterface
	Handlers     Handlers
	AuthProvider interfaces.AuthProviderInterface
}

type Handlers struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	GetList appnews.GetListHandler
}

type Commands struct {
}
