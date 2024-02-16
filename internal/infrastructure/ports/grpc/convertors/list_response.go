package convertors

import (
	"grpc/internal/domain/news"
	pg "grpc/internal/infrastructure/ports/grpc/proto_gen/grpc"
)

type ListResponseConvertorInterface interface {
	Convert([]news.New) *pg.NewsList
}
type ListResponseConvertor struct {
}

func (l ListResponseConvertor) Convert(list []news.New) *pg.NewsList {
	var newDto *pg.New
	var newsList []*pg.New

	for _, item := range list {
		newDto = &pg.New{
			Id:           &pg.UUID{Id: item.Id.String()},
			Title:        item.Title,
			Text:         item.Text,
			ActivePeriod: "",
			Status:       pg.Status.Enum(pg.Status(item.Status)),
			Media:        nil,
			CreatedAt:    nil,
			UpdatedAt:    nil,
			Tags:         nil,
		}

		newsList = append(newsList, newDto)
	}

	return &pg.NewsList{News: newsList}
}
