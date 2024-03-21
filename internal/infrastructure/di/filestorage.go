package di

import (
	minio "github.com/minio/minio-go"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"grpc/internal/infrastructure/adapters/storage/file"
	"log"
)

var FileStorage = []di.Def{
	{
		Name:  "FileStorageProvider",
		Scope: di.App,
		Build: func(c di.Container) (interface{}, error) {
			config := c.Get("ConfigProvider").(*viper.Viper)
			client, err := minio.New(
				config.GetString("MINIO_HOST"),
				config.GetString("MINIO_ACCESS_KEY"),
				config.GetString("MINIO_SECRET_KEY"),
				false)
			if err != nil {
				log.Fatalln(err)
			}

			return file.MinioStorageAdapter{
				Client: client,
			}, nil
		},
	},
}
