package application

import (
	appnewscommands "grpc/internal/application/news/commands"
	appnews "grpc/internal/application/news/queries"
	"grpc/internal/application/storage"
	appusers "grpc/internal/application/users/queries"
	"grpc/internal/domain/repository"
	"grpc/internal/shared/interfaces"
)

type Services struct {
	//Handler    appnews.ListHandler
	NewsRepository      repository.NewsRepositoryInterface
	TagsRepository      repository.TagsRepositoryInterface
	Transactor          storage.TransactorInterface
	Handlers            Handlers
	AuthProvider        interfaces.AuthProviderInterface
	FileStorageProvider interfaces.FileStorageProviderInterface
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
	Create appnewscommands.CreateHandlerInterface
}
