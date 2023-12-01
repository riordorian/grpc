package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib"
	"grpc/internal/domain/news"
)

// TODO: Code of connect/close methods duplicated on each repository

// Идея такая - на уровне адаптера создаю новую структуру транзактора, в которой будет осуществляться подключение к бд, закрытие соединения и указание на использование транзакций
// Возвращаться будет

type NewsRepository struct {
	db *Db
}

func (r NewsRepository) GetList(ctx context.Context) ([]news.New, error) {

	fmt.Println(123)

	return []news.New{}, nil

	/*var movies []interface{}
	if err := s.dbx.SelectContext(
		ctx,
		&movies,
		`SELECT * FROM forms`); err != nil {
		fmt.Println(movies)
	}*/

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
