package application

import (
	appnews "grpc/internal/application/news/queries"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
)

type Services struct {
	//Handler    appnews.ListHandler
	Repository news.RepositoryInterface
	Transactor storage.TransactorInterface
	Handlers   Handlers
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

func GetServices(repository news.RepositoryInterface, transactor storage.TransactorInterface) Services {
	return Services{
		/*Handler: appnews.ListHandler{
			Repo:       repository,
			Transactor: transactor,
		},*/

		Handlers: Handlers{
			Queries: Queries{
				GetList: appnews.NewGetListHandler(repository, transactor),
			},
			Commands: Commands{},
		},
	}
}
