package main

import (
	"context"
	"github.com/sarulabs/di"
	"grpc/dependencies"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports"
	"grpc/internal/shared/interfaces"
	"log"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	handlers := app.Get("ApplicationHandlers").(application.Handlers)
	configProvider := app.Get("ConfigProvider").(interfaces.ConfigProviderInterface)
	portsServices := ports.GetServices(configProvider, handlers)

	// TODO: graceful shutdown for http server on SIG or panic
	//if err := portsServices.HttpServer.RegisterHandlers(); err != nil {
	//	// TODO: How to operate panic?
	//	panic(err.Error())
	//}
	portsServices.GrpcServer.Serve()
	//go portsServices.HttpServer.ListenAndServe()

}
