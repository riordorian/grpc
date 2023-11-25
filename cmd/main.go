package main

import (
	"grpc/internal/infrastructure/ports"
)

func main() {

	interfaceServices := ports.GetServices()
	interfaceServices.GrpcServer.Serve()
	interfaceServices.GrpcServer.Server.GetServiceInfo()
}
