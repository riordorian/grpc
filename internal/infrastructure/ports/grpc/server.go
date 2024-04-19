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
	ListRequest   convertors.ListRequestConvertorInterface
	ListResponse  convertors.ListResponseConvertorInterface
	LoginRequest  convertors.UserLoginRequestConvertorInterface
	LoginResponse convertors.UserLoginResponseConvertorInterface
}

type NewsServer struct {
	grpc.UnimplementedNewsServer
	grpc.UnimplementedUsersServer
	Handlers   application.Handlers
	Server     *gp.Server
	Listener   net.Listener
	Container  di.Container
	Convertors Convertors
}

func (NewsServer) mustEmbedUnimplementedNewsServer()  {}
func (NewsServer) mustEmbedUnimplementedUsersServer() {}

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
