package postgres

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib"
	"grpc/internal/domain/news"
	"strings"
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

	query := map[string]interface{}{
		"sort":   req.Sort,
		"status": req.Status,
	}

	author := req.Author
	switch author {
	case uuid.Nil:
		query["author"] = "not null"
	default:
		// TODO: Move stringify to serializer
		query["author"] = author.String()
	}

	queryString := strings.Join(
		[]string{
			"SELECT * FROM news WHERE created_by=:author AND status=:status",
			"ORDER BY created_at",
			req.Sort,
		},
		" ")

	rows, err := r.Db.NamedQuery(ctx, queryString, query)

	if err != nil {
		return nil, err
	}

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

func (r NewsRepository) Insert(ctx context.Context, fields news.New) (uuid.UUID, error) {
	sql := "INSERT INTO news (title, text, created_by, status) VALUES (:title, :text, :created_by, :status)"
	_, err := r.Db.NamedExec(ctx, sql, fields)
	if err != nil {
		return uuid.New(), err
	}
	return uuid.New(), nil
}

func (NewsRepository) Update(id uuid.UUID, fields news.New) (bool, error) {
	return false, nil
}

func (NewsRepository) Delete(id uuid.UUID) (bool, error) {
	return false, nil
}
