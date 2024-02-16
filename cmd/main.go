package main

import (
	"context"
	"github.com/sarulabs/di"
	"grpc/dependencies"
	"grpc/internal/infrastructure/ports"
	"log"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	builder, builderErr := di.NewBuilder()
	if builderErr != nil {
		panic(builderErr.Error())
	}
	c := len(dependencies.ConfigServices) +
		len(dependencies.ConfigServices) +
		len(dependencies.RepositoryServices) +
		len(dependencies.AuthServices) +
		len(dependencies.ApplicationServices) +
		len(dependencies.PortsServices)

	var defs = make([]di.Def, 0, c)
	defs = append(defs, dependencies.ConfigServices...)
	defs = append(defs, dependencies.RepositoryServices...)
	defs = append(defs, dependencies.AuthServices...)
	defs = append(defs, dependencies.ApplicationServices...)
	defs = append(defs, dependencies.PortsServices...)
	builderErr = builder.Add(defs...)
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

	portsServices := app.Get("PortsServices").(ports.Services)

	// TODO: graceful shutdown for http server on SIG or panic
	//if err := portsServices.HttpServer.RegisterHandlers(); err != nil {
	//	// TODO: How to operate panic?
	//	panic(err.Error())
	//}
	portsServices.GrpcServer.Serve()
	//go portsServices.HttpServer.ListenAndServe()

}
