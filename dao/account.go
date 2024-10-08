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
		"INSERT INTO accounts (email,username, password_hash, icon_url, created_at) VALUES ($1, $2, $3, $4,$5)",
		account.Email,
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

func (a *Account) Login(ctx context.Context, tx *sqlx.Tx, account *object.Account) (*object.Account, error) {
	var acc object.Account
	err := a.db.Get(&acc, "SELECT * FROM accounts WHERE email = $1", account.Email)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}
