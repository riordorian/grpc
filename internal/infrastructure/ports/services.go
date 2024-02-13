package ports

import (
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports/grpc"
	"grpc/internal/infrastructure/ports/http"
	"grpc/internal/shared/interfaces"
)

type Services struct {
	GrpcServer *grpc.NewsServer
	HttpServer *http.Server
}

func GetServices(configProvider interfaces.ConfigProviderInterface, handlers application.Handlers) Services {
	return Services{
		GrpcServer: grpc.NewServer(configProvider, handlers),
		//HttpServer: http.GetServer(appServices.Handler),
	}
}
