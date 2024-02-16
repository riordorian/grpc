package dependencies

import (
	"context"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"grpc/internal/infrastructure/adapters/storage/postgres"
	"log"
)

var RepositoryServices = []di.Def{
	{
		Name:  "Database",
		Scope: di.App,
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
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return postgres.NewsRepository{
				Db: ctn.Get("Database").(*postgres.Db),
			}, nil
		},
	},
}
