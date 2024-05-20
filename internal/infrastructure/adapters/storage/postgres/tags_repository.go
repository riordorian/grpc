package postgres

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib"
	"grpc/internal/infrastructure/db"
)

// Идея такая - на уровне адаптера создаю новую структуру транзактора, в которой будет осуществляться подключение к бд, закрытие соединения и указание на использование транзакций
// Возвращаться будет

type TagsRepository struct {
	Db *db.Db
}

func (t TagsRepository) AddNewsTags(ctx context.Context, newsId uuid.UUID, userId uuid.UUID, tags []string) error {
	tagsIds := make([]uuid.UUID, 0, len(tags))
	var uid uuid.UUID
	for _, tag := range tags {
		// TODO: check duplicates by two fields user_id and text
		sql := "INSERT INTO tags (text, user_id) VALUES (:text, :user_id) ON CONFLICT DO NOTHING RETURNING id"
		preparedQuery, err := t.Db.PrepareNamed(sql)
		if err != nil {
			return err
		}

		err = preparedQuery.Get(&uid, map[string]interface{}{
			"text":    tag,
			"user_id": userId,
		})
		if err != nil {
			return err
		}
		tagsIds = append(tagsIds, uid)
	}

	for _, tagId := range tagsIds {
		sql := "INSERT INTO news_tags (new_id, tag_id) VALUES (:new_id, :tag_id) ON CONFLICT DO NOTHING"

		_, err := t.Db.NamedExec(ctx, sql, map[string]interface{}{
			"new_id": newsId,
			"tag_id": tagId,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (TagsRepository) Delete(id uuid.UUID) (bool, error) {
	return false, nil
}
