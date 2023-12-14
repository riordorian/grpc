package postgres

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib"
	"grpc/internal/domain/news"
)

// TODO: Code of connect/close methods duplicated on each repository

// Идея такая - на уровне адаптера создаю новую структуру транзактора, в которой будет осуществляться подключение к бд, закрытие соединения и указание на использование транзакций
// Возвращаться будет

type NewsRepository struct {
	Db *Db
}

// TODO: Is it correct to use structure from Domain layer

func (r NewsRepository) GetList(ctx context.Context, req news.ListRequest) ([]news.New, error) {
	var result []news.New
	dbx, err := GetDb(ctx)
	if err != nil {
		return nil, err
	}
	rows, _ := dbx.Model(ctx).Queryx("SELECT * FROM news")

	a := dbx.Model(ctx).NamedExec("SELECT * FROM news ")

	var newItem news.New
	for rows.Next() {
		if errScan := rows.StructScan(&newItem); err != nil {
			return nil, errScan
		} else {
			result = append(result, newItem)
		}
	}

	return result, nil
}

func (NewsRepository) GetById(uuid uuid.UUID) (news.New, error) {
	return news.New{}, nil
}

func (NewsRepository) Insert(news.New) (uuid.UUID, error) {
	return uuid.New(), nil
}

func (NewsRepository) Update(id uuid.UUID, fields news.New) (bool, error) {
	return false, nil
}

func (NewsRepository) Delete(id uuid.UUID) (bool, error) {
	return false, nil
}

func GetNewsRepository() NewsRepository {
	return NewsRepository{}
}
