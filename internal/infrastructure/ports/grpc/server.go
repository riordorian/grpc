package grpc

import (
	"context"
	"fmt"
	gp "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc/internal/infrastructure/adapters/storage/postgres"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
	"log"
	"net"
)

type Server struct {
	Server   *gp.Server
	Listener net.Listener
}

func NewServer() *Server {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	// TODO: Add logger
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := gp.NewServer()
	s := &Server{
		Server:   server,
		Listener: lis,
	}

	reflection.Register(s.Server)
	pg.RegisterNewsServer(s.Server, NewsServer{})
	//pg.RegisterPromosServer(server, pg.UnimplementedPromosServer{})

	return s
}

func (s *Server) Serve() {
	// TODO: Add logger
	if err := s.Server.Serve(s.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Serving...")
}

type NewsServer struct {
	pg.UnimplementedNewsServer
}

func (NewsServer) List(ctx context.Context, e *emptypb.Empty) (*emptypb.Empty, error) {
	dbx := postgres.GetDb()
	if err := dbx.Connect(ctx); err != nil {
		fmt.Println(err.Error())
	}
	defer func() {
		err := dbx.Close()
		fmt.Println(err.Error())
	}()

	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (NewsServer) mustEmbedUnimplementedNewsServer() {}
