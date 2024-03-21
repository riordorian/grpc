package di

import (
	"github.com/sarulabs/di"
	"grpc/internal/application"
	appnewscommands "grpc/internal/application/news/commands"
	appnews "grpc/internal/application/news/queries"
	appusers "grpc/internal/application/users/queries"
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/db"
	"grpc/internal/shared/interfaces"
)

var ApplicationServices = []di.Def{
	{
		Name:  "ApplicationHandlers",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			transactor := ctn.Get("Database").(*db.Db)
			repository := ctn.Get("NewsRepository").(news.RepositoryInterface)
			authProvider := ctn.Get("AuthProvider").(interfaces.AuthProviderInterface)
			fileStorageProvider := ctn.Get("FileStorageProvider").(interfaces.FileStorageProviderInterface)

			return application.Handlers{
				Queries: application.Queries{
					GetList: appnews.NewGetListHandler(repository, transactor),
					Login:   appusers.NewLoginHandler(authProvider),
				},
				Commands: application.Commands{
					Create: appnewscommands.NewCreateHandler(repository, transactor, fileStorageProvider),
				},
			}, nil
		},
	},
	{
		Name:  "AdaptersServices",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			dbx := ctn.Get("Database").(*db.Db)

			return adapters.Services{
				Database:            dbx,
				NewsRepository:      ctn.Get("NewsRepository").(news.RepositoryInterface),
				AuthProvider:        ctn.Get("AuthProvider").(interfaces.AuthProviderInterface),
				FileStorageProvider: ctn.Get("FileStorageProvider").(interfaces.FileStorageProviderInterface),
			}, nil
		},
	},
}
