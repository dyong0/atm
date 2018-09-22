package account

import "github.com/dyong0/atm/pkg/currency"

type Account struct {
	balance currency.Amount
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
		return newAmount, err
	}

	return amount, nil
}

func (a *Account) Total() uint32 {
	return a.balance.Total()
}

func NewAccount(currenyKind currency.CurrencyKind) *Account {
	amount, _ := currency.New(currenyKind, 0)

	return &Account{
		balance: amount,
	}
}
