package ports

import (
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports/grpc"
	"grpc/internal/infrastructure/ports/http"
)

type Services struct {
	GrpcServer *grpc.NewsServer
	HttpServer *http.Server
}

func GetServices(appServices application.Services) Services {
	return Services{
		GrpcServer: grpc.NewServer(appServices.Handlers),
		//HttpServer: http.GetServer(appServices.Handler),
	}
}
