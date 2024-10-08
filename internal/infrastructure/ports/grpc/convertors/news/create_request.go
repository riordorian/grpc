package news

import (
	"grpc/internal/shared/dto"
	pg "grpc/pkg/proto_gen/grpc"
)

type CreateRequestConvertorInterface interface {
	Convert(request *pg.CreateRequest) (dto.CreateRequest, error)
}

type CreateRequestConvertor struct {
}

func (l CreateRequestConvertor) Convert(req *pg.CreateRequest) (dto.CreateRequest, error) {

	createRequest := dto.CreateRequest{
		Title: req.GetTitle(),
		Text:  req.GetText(),
		Tags:  req.GetTags(),
		Media: req.GetMedia(),
	}

	return createRequest, nil
}
