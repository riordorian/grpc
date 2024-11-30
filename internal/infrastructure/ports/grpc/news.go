package grpc

import (
	"context"
	"grpc/internal/shared/dto"
	"grpc/pkg/proto_gen/grpc"
	"io"
	"log"
)

func (NewsServer) mustEmbedUnimplementedNewsServer() {}
func (s NewsServer) List(ctx context.Context, req *grpc.ListRequest) (*grpc.NewsList, error) {
	requestDto := new(dto.ListRequest)
	_, err := s.Convertors.Request.Convert(req, requestDto)
	if err != nil {
		return nil, err
	}

	list, listErr := s.Handlers.Queries.GetList.Handle(ctx, *requestDto)
	if listErr != nil {
		return nil, listErr
	}

	var newDto *grpc.New
	var newsList []*grpc.New

	for _, item := range list {
		newDto = &grpc.New{
			Id:           &grpc.UUID{Id: item.Id.String()},
			Title:        item.Title,
			Text:         item.Text,
			ActivePeriod: "",
			Status:       grpc.Status.Enum(grpc.Status(item.Status)),
			Media:        nil,
			CreatedAt:    nil,
			UpdatedAt:    nil,
			Tags:         nil,
		}

		newsList = append(newsList, newDto)
	}

	return &grpc.NewsList{News: newsList}, nil
}

func (s NewsServer) Create(stream grpc.News_CreateServer) error {
	requestDto := new(dto.CreateRequest)
	req, err := stream.Recv()
	_, err = s.Convertors.Request.Convert(req, requestDto)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		log.Fatal(err.Error())
	}

	createErr := s.Handlers.Commands.Create.Handle(stream.Context(), *requestDto)
	if createErr != nil {
		return createErr
	}

	return stream.SendAndClose(&grpc.CreateResponse{Status: true})
}
