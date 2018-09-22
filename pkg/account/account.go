package account

import "github.com/dyong0/atm/pkg/currency"

type Account struct {
	amount currency.Amount
}

func (a Account) Balance() uint32 {
	return a.amount.Total()
}

func (a *Account) Deposit(amount currency.Amount) error {
	newAmount, err := a.amount.Add(amount)
	if err != nil {
		return err
	}

	a.amount = newAmount

	return nil
}

func (a *Account) Withdraw(amount currency.Amount) (currency.Amount, error) {
	newAmount, err := a.amount.Subtract(amount)
	if err != nil {
		return newAmount, err
	}

	return amount, nil
}

func (a *Account) Total() uint32 {
	return a.amount.Total()
}

func New(currenyKind currency.CurrencyKind) *Account {
	amount, _ := currency.New(currenyKind, 0)

	return &Account{
		amount: amount,
	}
}
