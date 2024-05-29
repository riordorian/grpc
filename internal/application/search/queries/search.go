package queries

import (
	"context"
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

	_, err := s.SearchProvider.Search("news", fieldName, queryString)

	if err != nil {
		return "", err
	}

	return "hello", nil
}
