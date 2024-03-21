package grpc

import (
	"context"
	"fmt"
	"github.com/sarulabs/di"
	gp "google.golang.org/grpc"
	"grpc/internal/application"
	"grpc/internal/infrastructure/ports/grpc/convertors"
	"grpc/pkg/proto_gen/grpc"
	"io"
	"log"
	"net"
	"os"
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

func (s NewsServer) List(ctx context.Context, req *grpc.ListRequest) (*grpc.NewsList, error) {
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

func (s NewsServer) Create(stream grpc.News_CreateServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}

		media := req.GetMedia()

		for i := range media {
			file, err := os.Create(media[i].GetFileName())
			if err != nil {
				fmt.Println("Error creating file:", err)
				return err
			}
			file.Write(media[i].GetChunk())
			//fmt.Println(media[i].GetChunk())
		}

	}

	_, createErr := s.Handlers.Commands.Create.Handle("created.png")
	if createErr != nil {
		return createErr
	}

	return stream.SendAndClose(&grpc.CreateResponse{Status: true})
}

func (s NewsServer) Login(ctx context.Context, req *grpc.UserLoginRequest) (*grpc.UserLoginResponse, error) {
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
