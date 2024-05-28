package search

import (
	"github.com/google/uuid"
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
}

type BaseSearchProviderInterface interface {
	CreateIndex(index string) (*Response, error)
	DeleteIndex(index []string) (*Response, error)
	IndexDocument(index string, uuid uuid.UUID, document interface{}) (*Response, error)
	DeleteDocument(index string, id uuid.UUID) (*Response, error)
	Search(index string, query string) (*Response, error)
	SearchById(index string, id uuid.UUID) (*Response, error)
}
