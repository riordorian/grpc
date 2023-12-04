package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	url string
	dbx *sqlx.DB
}

const driverName = "pgx"

/*
	type Dbx interface {
		Begin() (Txx, error)
		Connect(ctx context.Context) error
		Close() error
	}

	type Txx interface {
		Commit() error
		Rollback() error
	}
*/
func GetDb() *Db {
	return &Db{
		url: "user=grpc password=password host=localhost port=5432 database=grpc sslmode=disable",
	}
}

func (d *Db) Connect(ctx context.Context) error {
	fmt.Println("DB Start")
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

func (d *Db) MakeTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	// begin transaction
	tx, err := d.dbx.Beginx()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	// run callback
	err = tFunc(addCtxTransact(ctx, tx))
	if err != nil {
		// if error, rollback
		if errRbck := tx.Rollback(); errRbck != nil {
			if errs := errors.Join(err, errRbck); errs != nil {
				panic("Errors during transaction")
			}
		}
		return err
	}

	// if no error, commit
	if errCmt := tx.Commit(); errCmt != nil {
		return errCmt
	}
	return nil
}

func (d *Db) Model(ctx context.Context) Querier {
	qa := &QueriesAdapter{}
	if tx := extractCtxTransact(ctx); tx != nil {
		return qa.adapt(d.dbx, tx)
	}

	return qa.adapt(d.dbx, nil)
}

// QueriesAdapter realization
type QueriesAdapter struct{}

type Querier interface {
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	NamedExec(query string, arg interface{}) (sql.Result, error)
}

func (q QueriesAdapter) adapt(dbx *sqlx.DB, tx *sqlx.Tx) Querier {
	if tx != nil {
		return tx
	}

	return dbx
}

/*func (d *Db) QueryRows(query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.dbx.Queryx(query, args)
}
func (d *Db) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return d.dbx.QueryRowx(query, args)
}*/
