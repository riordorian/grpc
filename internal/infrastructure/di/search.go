package di

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sarulabs/di"
)

var SearchProvider = []di.Def{
	{
		Name:  "SearchProvider",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			client, err := elasticsearch.NewClient(elasticsearch.Config{
				Addresses: []string{"http://elasticsearch:9200"},
			})

			return client, err
		},
	},
}
