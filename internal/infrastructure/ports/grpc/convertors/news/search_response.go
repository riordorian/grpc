package news

import (
	"grpc/internal/shared/dto"
	"grpc/pkg/proto_gen/grpc"
)

type SearchResponseConvertorInterface interface {
	Convert(searchResponse dto.SearchResponse) *grpc.SearchResponse
}
type SearchResponseConvertor struct {
}

func (l SearchResponseConvertor) Convert(searchResponse dto.SearchResponse) *grpc.SearchResponse {
	return &grpc.SearchResponse{Result: searchResponse.Body}
}
