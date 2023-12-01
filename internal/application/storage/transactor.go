package storage

import "context"

type Transactor interface {
	MakeTransaction(context.Context, func(ctx context.Context) error) error
}
