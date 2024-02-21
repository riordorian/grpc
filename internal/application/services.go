package application

import (
	appnews "grpc/internal/application/news/queries"
	"grpc/internal/application/storage"
	appusers "grpc/internal/application/users/queries"
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
	GetList appnews.GetListHandlerInterface
	Login   appusers.LoginHandlerInterface
}

type Commands struct {
}
