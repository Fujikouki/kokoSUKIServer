package usecase

import (
	"awesomeProject1/domain/object"
	"awesomeProject1/domain/repository"
	"context"
	"github.com/jmoiron/sqlx"
)

type AccountUsecase interface {
	CreateAccount(ctx context.Context, email, username, password, iconUrl string) error
}

type AccountR struct {
	db          *sqlx.DB
	accountRepo repository.AccountRepository
}

var _ AccountUsecase = (*AccountR)(nil)

func NewAccountUsecase(db *sqlx.DB, accountRepo repository.AccountRepository) *AccountR {
	return &AccountR{db: db, accountRepo: accountRepo}
}

func (a *AccountR) CreateAccount(ctx context.Context, email, username, password, iconUrl string) error {

	acc, err := object.NewAccount(email, username, password, iconUrl)
	if err != nil {
		return err
	}

	tx, err := a.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err := recover(); err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}

		err := tx.Commit()
		if err != nil {
			return
		}
	}()

	if err := a.accountRepo.Create(ctx, tx, acc); err != nil {
		return err
	}

	return nil
}
