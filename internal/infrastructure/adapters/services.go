package adapters

import (
	"grpc/internal/infrastructure/adapters/storage/postgres"
)

type Services struct {
	NewsRepository *postgres.NewsRepository
}

func GetServices() Services {
	return Services{
		//TODO: add url
		//NewsRepository: postgres.GetNewsRepository("postgresql://med_stat:Q}06Avm#Vbit@45.12.65.163:5432/med_stat"),
		//NewsRepository: postgres.GetNewsRepository("user=med_stat password=Q}06Avm#Vbit host=45.12.65.163 port=5432 database=med_stat sslmode=disable"),
		NewsRepository: postgres.GetNewsRepository("user=grpc password=password host=localhost port=5432 database=grpc sslmode=disable"),
	}
}
