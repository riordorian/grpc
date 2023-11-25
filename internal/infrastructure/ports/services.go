package ports

import (
	"grpc/internal/infrastructure/ports/grpc"
)

type Services struct {
	GrpcServer *grpc.Server
}

func GetServices() Services {
	return Services{GrpcServer: grpc.NewServer()}
}
