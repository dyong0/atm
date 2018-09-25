package account

import (
	"errors"
)

var (
	ErrAccountNotFound    = errors.New("Account not found")
	ErrInvalidAccount     = errors.New("Invalid account")
	ErrIDPasswordMismatch = errors.New("ID and password don't match")
)

type Repository interface {
	Account(id string) (*Account, error)
	VerifyAccount(account *Account, pw string) error
}

type repository struct {
	Repository
}

func NewRepository() (Repository, error) {
	return &repository{}, nil
}
