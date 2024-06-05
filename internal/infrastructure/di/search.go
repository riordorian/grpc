package di

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"grpc/internal/infrastructure/adapters/search"
)

var SearchServices = []di.Def{
	{
		Name:  "SearchProvider",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			config := c.Get("ConfigProvider").(*viper.Viper)
			dsn := config.GetString("ELASTIC_DSN")
			client, err := elasticsearch.NewClient(elasticsearch.Config{
				Addresses: []string{dsn},
			})

			if err != nil {
				return nil, err
			}

			return &search.ElasticAdapter{
				Client: client,
			}, err
		},
	},
}
