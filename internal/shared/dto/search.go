package dto

import (
	"net/http"
)

type SearchResponse struct {
	StatusCode int
	Header     http.Header
	Body       string
}
