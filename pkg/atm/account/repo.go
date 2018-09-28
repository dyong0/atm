package account

import (
	"errors"
)

var (
	ErrAccountNotFound    = errors.New("account not found")
	ErrIDPasswordMismatch = errors.New("id and password don't match")
)

type Repository interface {
	Account(id string) (*Account, error)
	Create(acc Account, holderName string, password string) error
	Update(account Account) error
	Delete(id string) error
}
