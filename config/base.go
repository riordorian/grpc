package config

import (
	"fmt"
	v "github.com/spf13/viper"
	"grpc/internal/shared/interfaces"
	"log"
)

func InitConfig() interfaces.ConfigProviderInterface {
	viper := v.GetViper()
	viper.AddConfigPath("../")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error while reading config file %s", err)
		// TODO: Panic handling()?
		panic(err.Error())
	}

	appEnv := viper.Get("APP_PROFILE")
	if appEnv == "local" {
		viper.Set("POSTGRES_HOST", viper.Get("POSTGRES_HOST_LOCAL"))
	}
	viper.SetDefault("POSTGRES_PORT", 5432)

	viper.SetDefault("GRPC_SERVER_HOST", "0.0.0.0")

	//dsn format: "user=grpc password=password host=localhost port=5432 database=grpc sslmode=disable",
	dbDsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d database=%s sslmode=disable",
		viper.Get("POSTGRES_USER"),
		viper.Get("POSTGRES_PASSWORD"),
		viper.Get("POSTGRES_HOST"),
		viper.GetInt32("POSTGRES_PORT"),
		viper.Get("POSTGRES_DB"),
	)
	viper.Set("POSTGRES_DSN", dbDsn)

	minioHost := viper.Get("MINIO_HOST")
	if appEnv == "local" {
		minioHost = viper.Get("MINIO_HOST_LOCAL")
	}

	minioUrl := fmt.Sprintf(
		"%s:%d", minioHost, viper.GetInt32("MINIO_PORT"),
	)
	viper.Set("MINIO_HOST", minioUrl)

	return viper
}
