package convertors

import (
	"github.com/google/uuid"
	"grpc/internal/domain/news"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
)

type ListRequestConvertorInterface interface {
	Convert(*pg.ListRequest) (news.ListRequest, error)
}

type ListRequestConvertor struct {
}

// TODO: Is it correct way to modify params here? Or should it be in the paramsRelolver like symfony?
func (l ListRequestConvertor) Convert(req *pg.ListRequest) (news.ListRequest, error) {
	page := *req.Page
	if page == 0 {
		page = 1
	}

	author, errParse := uuid.Parse(req.Author.GetId())
	if errParse != nil {
		return news.ListRequest{}, errParse
	}

	listRequest := news.ListRequest{
		// TODO: Is it correct way?
		Sort:   req.Sort.String(),
		Status: news.Status(req.Status.Number()),
		Query:  *req.Query,
		Page:   page,
		Author: author,
	}

	return listRequest, nil
}
