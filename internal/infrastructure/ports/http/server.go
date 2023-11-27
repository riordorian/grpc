package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	gp "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
	"net/http"
	"strings"
)

type Server struct {
	mux    *runtime.ServeMux
	ctx    context.Context
	opts   []gp.DialOption
	cancel context.CancelFunc
}

func GetServer() *Server {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	//defer cancel()
	mux := runtime.NewServeMux()
	opts := []gp.DialOption{gp.WithTransportCredentials(insecure.NewCredentials())}

	return &Server{
		mux:    mux,
		ctx:    ctx,
		opts:   opts,
		cancel: cancel,
	}
}

func (s *Server) RegisterHandlers() error {
	var err []string
	if er := grpc.RegisterNewsHandlerFromEndpoint(s.ctx, s.mux, "localhost:50051", s.opts); er != nil {
		err = append(err, er.Error())
	}

	if er2 := grpc.RegisterPromosHandlerFromEndpoint(s.ctx, s.mux, "localhost:50051", s.opts); er2 != nil {
		err = append(err, er2.Error())
	}

	if len(err) != 0 {
		return errors.New(strings.Join(err, ","))
	}

	return nil
}

func (s *Server) ListenAndServe() {
	defer s.cancel()
	fmt.Println("server listening at 8081")
	if err := http.ListenAndServe(":8081", s.mux); err != nil {
		panic(err)
	}

	fmt.Println("123")
}
