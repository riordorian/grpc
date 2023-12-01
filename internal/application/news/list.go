package news

import (
	"grpc/internal/application/storage"
	"grpc/internal/domain/news"
)

type ListHandler struct {
	repo       news.Repository
	transactor storage.Transactor
}
