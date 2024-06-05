package di

import (
	"github.com/sarulabs/di"
	"grpc/internal/application"
	appnewscommands "grpc/internal/application/news/commands"
	appnews "grpc/internal/application/news/queries"
	appsearch "grpc/internal/application/search/queries"
	appusers "grpc/internal/application/users/queries"
	"grpc/internal/domain/repository"
	"grpc/internal/domain/search"
	"grpc/internal/infrastructure/db"
	"grpc/internal/shared/interfaces"
)

var ApplicationServices = []di.Def{
	{
		Name:  "ApplicationHandlers",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			transactor := ctn.Get("Database").(*db.Db)
			newsRepository := ctn.Get("NewsRepository").(repository.NewsRepositoryInterface)
			tagsRepository := ctn.Get("TagsRepository").(repository.TagsRepositoryInterface)
			authProvider := ctn.Get("AuthProvider").(interfaces.AuthProviderInterface)
			fileStorageProvider := ctn.Get("FileStorageProvider").(interfaces.FileStorageProviderInterface)
			searchProvider := ctn.Get("SearchProvider").(search.BaseSearchProviderInterface)

			return application.Handlers{
				Queries: application.Queries{
					GetList: appnews.NewGetListHandler(newsRepository, transactor),
					Search:  appsearch.NewSearchHandler(searchProvider),
					Login:   appusers.NewLoginHandler(authProvider),
				},
				Commands: application.Commands{
					Create: appnewscommands.NewCreateHandler(newsRepository, tagsRepository, transactor, fileStorageProvider),
				},
			}, nil
		},
	},
}
