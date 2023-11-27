package postgres

import (
	"context"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// TODO: Code of connect/close methods duplicated on each repository

const driverName = "pgx"

type NewsRepository struct {
	databaseUrl string
	dbx         *sqlx.DB
}

func GetNewsRepository(databaseUrl string) *NewsRepository {
	return &NewsRepository{
		databaseUrl: databaseUrl,
	}
}

func (s *NewsRepository) connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, s.databaseUrl)
	if err != nil {
		return err
	}

	s.dbx = dbx
	return nil
}

func (s *NewsRepository) close() error {
	return s.dbx.Close()
}
