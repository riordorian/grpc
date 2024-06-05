package grpc

import (
	"fmt"
	"github.com/sarulabs/di"
	gp "google.golang.org/grpc"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports/grpc/convertors"
	"grpc/pkg/proto_gen/grpc"
	"log"
	"net"
)

type Convertors struct {
	ListRequest    convertors.ListRequestConvertorInterface
	ListResponse   convertors.ListResponseConvertorInterface
	CreateRequest  convertors.CreateRequestConvertorInterface
	LoginRequest   convertors.UserLoginRequestConvertorInterface
	LoginResponse  convertors.UserLoginResponseConvertorInterface
	SearchResponse convertors.SearchResponseConvertorInterface
}

type NewsServer struct {
	grpc.UnimplementedNewsServer
	grpc.UnimplementedUsersServer
	grpc.UnimplementedSearchServer
	Handlers   application.Handlers
	Server     *gp.Server
	Listener   net.Listener
	Container  di.Container
	Convertors Convertors
}

func (s NewsServer) SetHandlers(handlers application.Handlers) {
	s.Handlers = handlers
}

func (s NewsServer) Serve() {
	// TODO: Add logger
	if err := s.Server.Serve(s.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Serving...")
}
