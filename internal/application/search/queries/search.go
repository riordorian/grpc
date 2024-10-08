package queries

import (
	"context"
	"fmt"
	"grpc/internal/domain/search"
)

type SearchHandler struct {
	SearchProvider search.BaseSearchProviderInterface
}

type SearchHandlerInterface interface {
	Handle(ctx context.Context, fieldName string, queryString string) (string, error)
}

func NewSearchHandler(s search.BaseSearchProviderInterface) SearchHandlerInterface {
	return SearchHandler{
		SearchProvider: s,
	}
}

func (s SearchHandler) Handle(ctx context.Context, fieldName string, queryString string) (string, error) {

	res, err := s.SearchProvider.Search("news", fieldName, queryString)

	fmt.Println(res)
	if err != nil {
		return "", err
	}

	return "hello", nil
}
