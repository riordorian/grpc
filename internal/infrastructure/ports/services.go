package ports

import (
	"grpc/internal/infrastructure/ports/grpc"
	"grpc/internal/infrastructure/ports/http"
)

type Services struct {
	GrpcServer *grpc.Server
	HttpServer *http.Server
}

func GetServices() Services {
	return Services{
		GrpcServer: grpc.NewServer(),
		HttpServer: http.GetServer(),
	}
}
