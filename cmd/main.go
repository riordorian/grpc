package main

import (
	"grpc/internal/application"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/ports"
)

func main() {
	adapters := adapters.GetServices()
	appServices := application.GetServices(adapters.NewsRepository, adapters.Database)
	portsServices := ports.GetServices(appServices)

	// TODO: graceful shutdown for http server on SIG or panic
	if err := portsServices.HttpServer.RegisterHandlers(); err != nil {
		// TODO: How to operate panic?
		panic(err.Error())
	}
	go portsServices.GrpcServer.Serve()
	portsServices.HttpServer.ListenAndServe()
}
