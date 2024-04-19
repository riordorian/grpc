package grpc

import (
	"context"
	"fmt"
	"grpc/pkg/proto_gen/grpc"
	"io"
	"os"
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
