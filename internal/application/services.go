package application

import (
	appnews "grpc/internal/application/news"
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
)

type Services struct {
	Handler    appnews.ListHandler
	Repository news.Repository
	Transactor storage.Transactor
}

func GetServices(repository news.Repository, transactor storage.Transactor) Services {
	return Services{
		Handler: appnews.ListHandler{
			Repo:       repository,
			Transactor: transactor,
		},
	}
}
