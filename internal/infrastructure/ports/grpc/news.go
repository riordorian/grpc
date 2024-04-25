package grpc

import (
	"context"
	"grpc/pkg/proto_gen/grpc"
	"io"
	"log"
)

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
	req, err := stream.Recv()
	if err == io.EOF {
		return nil
	}

	//createRequest, err := s.Convertors.CreateRequest.Convert(req)
	_, err = s.Convertors.CreateRequest.Convert(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, createErr := s.Handlers.Commands.Create.Handle("created.png")
	if createErr != nil {
		return createErr
	}

	return stream.SendAndClose(&grpc.CreateResponse{Status: true})
}
