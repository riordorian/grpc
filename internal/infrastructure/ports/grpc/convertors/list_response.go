package convertors

import (
	"grpc/internal/domain/news"
	"grpc/pkg/proto_gen/grpc"
)

type ListResponseConvertorInterface interface {
	Convert([]news.New) *grpc.NewsList
}
type ListResponseConvertor struct {
}

func (l ListResponseConvertor) Convert(list []news.New) *grpc.NewsList {
	var newDto *grpc.New
	var newsList []*grpc.New

	for _, item := range list {
		newDto = &grpc.New{
			Id:           &grpc.UUID{Id: item.Id.String()},
			Title:        item.Title,
			Text:         item.Text,
			ActivePeriod: "",
			Status:       grpc.Status.Enum(grpc.Status(item.Status)),
			Media:        nil,
			CreatedAt:    nil,
			UpdatedAt:    nil,
			Tags:         nil,
		}

		newsList = append(newsList, newDto)
	}

	return &grpc.NewsList{News: newsList}
}
