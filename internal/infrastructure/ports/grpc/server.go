package grpc

import (
	"fmt"
	gp "google.golang.org/grpc"
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
	s := &Server{
		Server:   gp.NewServer(),
		Listener: lis,
	}
	return s
}

func (s *Server) Serve() {
	// TODO: Add logger
	if err := s.Server.Serve(s.Listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Serving...")
}
