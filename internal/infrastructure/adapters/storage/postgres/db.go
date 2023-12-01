package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	url string
	dbx *sqlx.DB
}

const driverName = "pgx"

func GetDb() *Db {
	return &Db{
		url: "user=grpc password=password host=localhost port=5432 database=grpc sslmode=disable",
	}
}

func (d *Db) Connect(ctx context.Context) error {
	dbx, err := sqlx.ConnectContext(ctx, driverName, d.url)

	if err != nil {
		return err
	}

	d.dbx = dbx
	return nil
}

func (d *Db) Close() error {
	return d.dbx.Close()
}
