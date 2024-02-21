package dependencies

import (
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"grpc/internal/infrastructure/adapters/auth"
	"net/http"
)

var AuthServices = []di.Def{
	{
		Name:  "AuthProvider",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			config := c.Get("ConfigProvider").(*viper.Viper)
			return auth.Keycloak{
				BaseUrl:    config.GetString("KEYCLOAK_BASE_URL"),
				Realm:      config.GetString("KEYCLOAK_REALM"),
				ClientId:   config.GetString("KEYCLOAK_CLIENT_ID"),
				GrantType:  config.GetString("KEYCLOAK_GRANT_TYPE"),
				HTTPClient: new(http.Client),
			}, nil
		},
	},
}
