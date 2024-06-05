package grpc

import (
	"context"
	"grpc/pkg/proto_gen/grpc"
)

func (NewsServer) mustEmbedUnimplementedSearchServer() {}

func (s NewsServer) SearchQuery(ctx context.Context, req *grpc.SearchRequest) (*grpc.SearchResponse, error) {
	resp, err := s.Handlers.Queries.Search.Handle(ctx, req.GetField(), req.GetQueryString())
	if err != nil {
		return nil, err
	}

	return &grpc.SearchResponse{Result: resp}, nil
}
