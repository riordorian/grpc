package dependencies

import (
	"github.com/sarulabs/di"
	"grpc/internal/application"
	appnews "grpc/internal/application/news/queries"
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/adapters/storage/postgres"
	"grpc/internal/shared/interfaces"
)

var ApplicationServices = []di.Def{
	{
		Name:  "ApplicationHandlers",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			transactor := ctn.Get("Database").(*postgres.Db)
			repository := ctn.Get("NewsRepository").(news.RepositoryInterface)

			return application.Handlers{
				Queries: application.Queries{
					GetList: appnews.NewGetListHandler(repository, transactor),
				},
				Commands: application.Commands{},
			}, nil
		},
	},
	{
		Name:  "AdaptersServices",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			dbx := ctn.Get("Database").(*postgres.Db)

			return adapters.Services{
				Database:       dbx,
				NewsRepository: ctn.Get("NewsRepository").(news.RepositoryInterface),
				AuthProvider:   ctn.Get("AuthProvider").(interfaces.AuthProviderInterface),
			}, nil
		},
	},
}
