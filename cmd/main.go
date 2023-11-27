package main

import (
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/ports"
)

func main() {

	_ = adapters.GetServices()

	portsServices := ports.GetServices()

	if err := portsServices.HttpServer.RegisterHandlers(); err != nil {
		// TODO: How to operate panic?
		panic(err.Error())
	}
	portsServices.HttpServer.ListenAndServe()
	portsServices.GrpcServer.Serve()
}
