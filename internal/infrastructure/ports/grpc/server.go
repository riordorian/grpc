package grpc

import (
	"context"
	"fmt"
	"github.com/sarulabs/di"
	gp "google.golang.org/grpc"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports/grpc/convertors"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
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
	pg.UnimplementedNewsServer
	pg.UnimplementedUsersServer
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

func (s NewsServer) List(ctx context.Context, req *pg.ListRequest) (*pg.NewsList, error) {
	listRequest, err := s.Convertors.ListRequest.Convert(req)
	if err != nil {
		return nil, err
	}

	list, listErr := s.Handlers.Queries.GetList.Handle(ctx, listRequest)
	if listErr != nil {
		return nil, listErr
	}

	return s.Convertors.ListResponse.Convert(list), nil
}

func (s NewsServer) Login(ctx context.Context, req *pg.UserLoginRequest) (*pg.UserLoginResponse, error) {
	loginRequest, err := s.Convertors.LoginRequest.Convert(req)
	if err != nil {
		return nil, err
	}

	token, err := s.Handlers.Queries.Login.Handle(ctx, loginRequest)
	if err != nil {
		return nil, err
	}
	rawToken, err := s.Convertors.LoginResponse.Convert(token)
	if err != nil {
		return nil, err

	}

	return rawToken, nil
}
