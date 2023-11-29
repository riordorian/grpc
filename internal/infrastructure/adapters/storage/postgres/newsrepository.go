package postgres

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// TODO: Code of connect/close methods duplicated on each repository

const driverName = "pgx"

type NewsRepository struct {
	databaseUrl string
	dbx         *sqlx.DB
}

func (s *NewsRepository) GetAll(ctx context.Context) {
	if err := s.connect(ctx); err != nil {
		fmt.Println(s.databaseUrl)
		fmt.Println(err.Error())
	}
	fmt.Println(123)

	defer s.close()

	/*var movies []interface{}
	if err := s.dbx.SelectContext(
		ctx,
		&movies,
		`SELECT * FROM forms`); err != nil {
		fmt.Println(movies)
	}*/

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
