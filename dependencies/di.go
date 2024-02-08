package dependencies

import (
	"context"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"grpc/config"
	"grpc/internal/domain/news"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/adapters/storage/postgres"
	"log"
)

var Services = []di.Def{
	{
		Name:  "ConfigProvider",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			return config.InitConfig(), nil
		},
	},
	{
		Name:  "Database",
		Scope: di.Request,
		Build: func(c di.Container) (interface{}, error) {
			config := c.Get("ConfigProvider").(*viper.Viper)
			ctx, err := postgres.GetContextDb(context.Background(), config.Get("POSTGRES_DSN").(string))
			if err != nil {
				log.Fatal(err.Error())
			}

			db, err := postgres.GetDb(ctx)
			if err != nil {
				return nil, err
			}
			return db, nil
		},
	},
	{
		Name:  "NewsRepository",
		Scope: di.Request,
		Build: func(ctn di.Container) (interface{}, error) {
			return postgres.GetNewsRepository(ctn.Get("Database").(*postgres.Db)), nil
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
			}, nil
		},
	},
}
