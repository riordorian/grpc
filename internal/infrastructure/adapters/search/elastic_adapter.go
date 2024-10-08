package search

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"grpc/internal/shared/dto"
	"io"
	"log"
)

type ElasticAdapter struct {
	Client *elasticsearch.Client
}

func (e ElasticAdapter) CreateIndex(index string) (*dto.SearchResponse, error) {
	res, err := e.Client.Indices.Create(index)

	return &dto.SearchResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       e.getBody(res.Body),
	}, err
}

func (e ElasticAdapter) DeleteIndex(index []string) (*dto.SearchResponse, error) {
	res, err := e.Client.Indices.Delete(index)

	return &dto.SearchResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       e.getBody(res.Body),
	}, err
}

func (e ElasticAdapter) IndexDocument(index string, uuid uuid.UUID, document interface{}) (*dto.SearchResponse, error) {
	doc, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}
	docId := uuid.String()
	res, err := e.Client.Index(index, bytes.NewReader(doc), e.Client.Index.WithDocumentID(docId))

	return &dto.SearchResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       e.getBody(res.Body),
	}, err
}

func (e ElasticAdapter) DeleteDocument(index string, id uuid.UUID) (*dto.SearchResponse, error) {
	res, err := e.Client.Delete(index, id.String())

	return &dto.SearchResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       e.getBody(res.Body),
	}, err
}

func (e ElasticAdapter) Search(index string, fieldName string, queryString string) (*dto.SearchResponse, error) {
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

	return &dto.SearchResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       e.getBody(res.Body),
	}, err
}

func (e ElasticAdapter) SearchById(index string, id uuid.UUID) (*dto.SearchResponse, error) {
	res, err := e.Client.Get(index, id.String())

	return &dto.SearchResponse{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		Body:       e.getBody(res.Body),
	}, err
}

func (e ElasticAdapter) getBody(body io.ReadCloser) string {
	defer func() {
		if err := body.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	resBody, err := io.ReadAll(body)
	if err != nil {
		return ""
	}

	return string(resBody)
}
