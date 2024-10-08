package search

import (
	"github.com/google/uuid"
	"grpc/internal/shared/dto"
)

type BaseSearchProviderInterface interface {
	CreateIndex(index string) (*dto.SearchResponse, error)
	DeleteIndex(index []string) (*dto.SearchResponse, error)
	IndexDocument(index string, uuid uuid.UUID, document interface{}) (*dto.SearchResponse, error)
	DeleteDocument(index string, id uuid.UUID) (*dto.SearchResponse, error)
	Search(index string, fieldName string, queryString string) (*dto.SearchResponse, error)
	SearchById(index string, id uuid.UUID) (*dto.SearchResponse, error)
}
