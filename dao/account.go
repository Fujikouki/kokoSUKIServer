package dao

import (
	"awesomeProject1/domain/object"
	"awesomeProject1/domain/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type (
	Account struct {
		db *sqlx.DB
	}
)

var _ repository.AccountRepository = (*Account)(nil)

func NewAccount(db *sqlx.DB) *Account {
	return &Account{db: db}
}

func (a *Account) Create(ctx context.Context, tx *sqlx.Tx, account *object.Account) error {
	_, err := a.db.Exec(
		"INSERT INTO accounts (username, password_hash, icon_url, created_at) VALUES ($1, $2, $3, $4)",
		account.Username,
		account.PasswordHash,
		account.IconUrl,
		account.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
