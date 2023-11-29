package main

import (
	context2 "context"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/ports"
)

func main() {

	adapters2 := adapters.GetServices()
	context := context2.Background()
	adapters2.NewsRepository.GetAll(context)

	portsServices := ports.GetServices()

	// TODO: graceful shutdown for http server
	/*if err := portsServices.HttpServer.RegisterHandlers(); err != nil {
		// TODO: How to operate panic?
		panic(err.Error())
	}*/
	//portsServices.HttpServer.ListenAndServe()
	portsServices.GrpcServer.Serve()
}
