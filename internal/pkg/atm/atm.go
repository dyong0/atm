package atm

import (
	"github.com/dyong0/atm/internal/pkg/atm/account"
	"github.com/dyong0/atm/internal/pkg/atm/account/method"
	"github.com/dyong0/atm/internal/pkg/atm/currency"
)

type ATM struct {
	accountRepo    account.Repository
	currentAccount *account.Account
}

func (a *ATM) ReadAccount(accMethod method.Method) error {
	acc, err := a.accountRepo.Account(accMethod.AccountID())
	if err != nil {
		return err
	}
	if err = acc.Authenticate(accMethod.Password()); err != nil {
		return err
	}

	a.currentAccount = acc

	return nil
}

func (a *ATM) Deposit(amount currency.Amount) error {
	return a.currentAccount.Deposit(amount)
}

func (a *ATM) Withdraw(amount currency.Amount) (currency.Amount, error) {
	return a.currentAccount.Withdraw(amount)
}

func (a *ATM) Balance() uint32 {
	return a.currentAccount.Balance()
}

func (a *ATM) CurrencyKind() currency.CurrencyKind {
	return a.currentAccount.CurrencyKind()
}

func (a *ATM) Close() error {
	a.currentAccount = nil
	return nil
}

func NewATM(accRepo account.Repository) (*ATM, error) {
	return &ATM{accountRepo: accRepo}, nil
}
