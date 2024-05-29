package search

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"grpc/internal/domain/search"
)

type ElasticAdapter struct {
	Client *elasticsearch.Client
}

func (e ElasticAdapter) CreateIndex(index string) (*search.Response, error) {
	res, err := e.Client.Indices.Create(index)

	return &search.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}, err
}

func (e ElasticAdapter) DeleteIndex(index []string) (*search.Response, error) {
	res, err := e.Client.Indices.Delete(index)

	return &search.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}, err
}

func (e ElasticAdapter) IndexDocument(index string, uuid uuid.UUID, document interface{}) (*search.Response, error) {
	doc, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	docId := uuid.String()
	res, err := e.Client.Index(index, bytes.NewReader(doc), e.Client.Index.WithDocumentID(docId))

	return &search.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}, err
}

func (e ElasticAdapter) DeleteDocument(index string, id uuid.UUID) (*search.Response, error) {
	res, err := e.Client.Delete(index, id.String())

	return &search.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}, err
}

func (e ElasticAdapter) Search(index string, fieldName string, queryString string) (*search.Response, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				fieldName: queryString,
			},
		},
	}
	queryBytes, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	res, err := e.Client.Search(
		e.Client.Search.WithIndex(index),
		e.Client.Search.WithBody(bytes.NewReader(queryBytes)),
	)

	return &search.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}, err
}

func (e ElasticAdapter) SearchById(index string, id uuid.UUID) (*search.Response, error) {
	res, err := e.Client.Get(index, id.String())

	return &search.Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       res.Body,
	}, err
}
