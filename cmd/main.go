package main

import (
	"context"
	"github.com/sarulabs/di"
	"grpc/dependencies"
	"grpc/internal/application"
	"grpc/internal/infrastructure/adapters"
	"grpc/internal/infrastructure/ports"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//config := config.InitConfig()

	builder, builderErr := di.NewBuilder()
	if builderErr != nil {
		panic(builderErr.Error())
	}

	builderErr = builder.Add(dependencies.Services...)
	if builderErr != nil {
		log.Fatal(builderErr.Error())
	}

	app := builder.Build()
	defer func() {
		diDeleteErr := app.Delete()
		if diDeleteErr != nil {
			panic(diDeleteErr.Error())
		}
	}()

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
