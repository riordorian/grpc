package grpc

import (
	"fmt"
	"github.com/sarulabs/di"
	gp "google.golang.org/grpc"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports/grpc/convertors/news"
	"grpc/internal/infrastructure/ports/grpc/convertors/users"
	"grpc/pkg/proto_gen/grpc"
	"log"
	"net"
)

type Convertors struct {
	ListRequest    news.ListRequestConvertorInterface
	ListResponse   news.ListResponseConvertorInterface
	CreateRequest  news.CreateRequestConvertorInterface
	LoginRequest   users.UserLoginRequestConvertorInterface
	LoginResponse  users.UserLoginResponseConvertorInterface
	SearchResponse news.SearchResponseConvertorInterface
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
