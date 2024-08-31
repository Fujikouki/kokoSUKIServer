package repository

import (
	"awesomeProject1/domain/object"
	"context"
	"github.com/jmoiron/sqlx"
)

type AccountRepository interface {
	Create(ctx context.Context, tx *sqlx.Tx, account *object.Account) error
}
