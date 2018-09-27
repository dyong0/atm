package account

import (
	"errors"
)

var (
	ErrAccountNotFound    = errors.New("Account not found")
	ErrIDPasswordMismatch = errors.New("ID and password don't match")
)

type Repository interface {
	Account(id string) (*Account, error)
	Create(acc Account, holderName string, password string) error
	Update(account Account) error
	Delete(id string) error
}
