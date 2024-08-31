package usecase

import (
	"awesomeProject1/domain/object"
	"awesomeProject1/domain/repository"
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type AccountUsecase interface {
	CreateAccount(ctx context.Context, email, username, password, iconUrl string) error
	Login(ctx context.Context, email, password string) (*object.Account, error)
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

func (a *AccountR) Login(ctx context.Context, email, password string) (*object.Account, error) {
	acc, err := a.accountRepo.Login(ctx, nil, &object.Account{Email: email, PasswordHash: password})
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.PasswordHash), []byte(password))
	if err != nil {
		// パスワードが一致しない場合の処理
		fmt.Println("Password does not match")
		return nil, err
	} else {
		// パスワードが一致する場合の処理
		fmt.Println("Password match")
		return acc, nil
	}
}
