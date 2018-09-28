package account

import (
	"errors"

	"github.com/dyong0/atm/pkg/atm/currency"
)

var (
	ErrInvalidAuthentication = errors.New("invalid authentication")
)

type Account struct {
	balance currency.Amount
	id      string
	pw      string
}

func (a Account) ID() string {
	return a.id
}
func (a Account) Password() string {
	return a.pw
}

func (a Account) Authenticate(password string) error {
	if a.pw != password {
		return ErrInvalidAuthentication
	}

	return nil
}

func (a Account) Balance() uint32 {
	return a.balance.Total()
}

func (a *Account) Deposit(amount currency.Amount) error {
	newAmount, err := a.balance.Add(amount)
	if err != nil {
		return err
	}

	a.balance = newAmount

	return nil
}

func (a *Account) Withdraw(amount currency.Amount) (currency.Amount, error) {
	newAmount, err := a.balance.Subtract(amount)
	if err != nil {
		return amount, err
	}

	a.balance = newAmount

	return amount, nil
}

func (a *Account) Total() uint32 {
	return a.balance.Total()
}

func (a *Account) CurrencyKind() currency.CurrencyKind {
	return a.balance.CurrencyKind()
}

func NewAccount(currenyKind currency.CurrencyKind) *Account {
	amount, _ := currency.NewAmount(currenyKind, 0)

	return &Account{
		balance: amount,
	}
}
