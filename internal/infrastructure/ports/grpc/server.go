package grpc

import (
	"context"
	"fmt"
	gp "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	appnews "grpc/internal/application/news"
	"grpc/internal/infrastructure/adapters/storage/postgres"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
	"log"
	"net"
)

type NewsServer struct {
	pg.UnimplementedNewsServer
	Handler  appnews.ListHandler
	Server   *gp.Server
	Listener net.Listener
}

func NewServer(handler appnews.ListHandler) *NewsServer {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	// TODO: Add logger
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := gp.NewServer()
	s := &NewsServer{
		Server:   server,
		Listener: lis,
		Handler:  handler,
	}

	reflection.Register(s.Server)
	pg.RegisterNewsServer(s.Server, NewsServer{})
	//pg.RegisterPromosServer(server, pg.UnimplementedPromosServer{})

	return s
}

func (s NewsServer) Serve() {
	// TODO: Add logger
	if err := s.Server.Serve(s.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Serving...")
}

func (s NewsServer) List(ctx context.Context, e *emptypb.Empty) (*emptypb.Empty, error) {

	dbx := postgres.GetDb()
	if err := dbx.Connect(ctx); err != nil {
		fmt.Println(err.Error())
	}

	eR := s.Handler.List(ctx)

	if eR != nil {
		fmt.Println(eR.Error())
	}

	defer func() {
		err := dbx.Close()
		fmt.Println(err.Error())
	}()

	return e, nil
}

func (NewsServer) mustEmbedUnimplementedNewsServer() {}
