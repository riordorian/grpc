package main

import (
	"context"
	"grpc/config"
	"grpc/internal/application"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/ports"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	config.InitConfig()

	adapters, err := adapters.GetServices(ctx)
	if err != nil {
		panic(err.Error())
	}
	appServices := application.GetServices(adapters.NewsRepository, adapters.Database)
	portsServices := ports.GetServices(appServices)

	// TODO: graceful shutdown for http server on SIG or panic
	//if err := portsServices.HttpServer.RegisterHandlers(); err != nil {
	//	// TODO: How to operate panic?
	//	panic(err.Error())
	//}
	portsServices.GrpcServer.Serve()
	//go portsServices.HttpServer.ListenAndServe()
}
