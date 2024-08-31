package object

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Account struct {
	ID           int       `db:"id"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	IconUrl      string    `db:"icon_url"`
	CreatedAt    time.Time `db:"created_at"`
}

func NewAccount(username, password, iconUrl string) (*Account, error) {
	newAccount := &Account{
		Username:     username,
		PasswordHash: password,
		IconUrl:      iconUrl,
		CreatedAt:    time.Now(),
	}

	if err := newAccount.SetPassword(password); err != nil {
		return nil, fmt.Errorf("set password error: %w", err)
	}

	return newAccount, nil
}

func (a *Account) SetPassword(pass string) error {
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", err)
	}
	return string(hash), nil
}
