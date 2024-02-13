package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sarulabs/di"
	gp "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc/internal/application"
	"grpc/internal/domain/news"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
	"grpc/internal/shared/interfaces"
	"log"
	"net"
)

type NewsServer struct {
	pg.UnimplementedNewsServer
	Handlers  application.Handlers
	Server    *gp.Server
	Listener  net.Listener
	Container di.Container
}

func (NewsServer) mustEmbedUnimplementedNewsServer() {}

func NewServer(configProvider interfaces.ConfigProviderInterface, handlers application.Handlers) *NewsServer {
	address := fmt.Sprintf("%s:%s", configProvider.GetString("GRPC_SERVER_HOST"), configProvider.GetString("GRPC_SERVER_PORT"))
	lis, err := net.Listen("tcp", address)
	// TODO: Add logger
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := gp.NewServer(gp.UnaryInterceptor(UnaryServerInterceptor))
	s := &NewsServer{
		Server:   server,
		Listener: lis,
		Handlers: handlers,
	}

	reflection.Register(s.Server)
	//pg.RegisterNewsServer(s.Server, NewsServer{})
	pg.RegisterNewsServer(s.Server, s)
	//pg.RegisterPromosServer(server, pg.UnimplementedPromosServer{})

	return s
}

func UnaryServerInterceptor(ctx context.Context, req any, info *gp.UnaryServerInfo, handler gp.UnaryHandler) (resp any, err error) {
	fmt.Println("UnaryServerInterceptor")

	return handler(ctx, req)
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

func (s NewsServer) List(ctx context.Context, req *pg.ListRequest) (*pg.NewsList, error) {

	// TODO: Move to list request serializer
	page := *req.Page
	if page == 0 {
		page = 1
	}

	author, errParse := uuid.Parse(req.Author.GetId())
	if errParse != nil {
		fmt.Println(errParse)
	}

	listRequest := news.ListRequest{
		// TODO: Is it correct way?
		Sort:   req.Sort.String(),
		Status: news.Status(req.Status.Number()),
		Query:  *req.Query,
		Page:   page,
		Author: author,
	}

	list, listErr := s.Handlers.Queries.GetList.Handle(ctx, listRequest)

	if listErr != nil {
		return nil, listErr
	}

	var newDto *pg.New
	var newsList []*pg.New

	for _, item := range list {
		newDto = &pg.New{
			Id:           &pg.UUID{Id: item.Id.String()},
			Title:        item.Title,
			Text:         item.Text,
			ActivePeriod: "",
			Status:       pg.Status.Enum(pg.Status(item.Status)),
			Media:        nil,
			CreatedAt:    nil,
			UpdatedAt:    nil,
			Tags:         nil,
		}

		newsList = append(newsList, newDto)
	}

	return &pg.NewsList{News: newsList}, nil
}
