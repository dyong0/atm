package atm

import (
	"github.com/dyong0/atm/internal/pkg/atm/account"
	"github.com/dyong0/atm/internal/pkg/atm/account/medium"
	"github.com/dyong0/atm/internal/pkg/atm/currency"
)

type ATM struct {
	accountRepo    account.Repository
	currentAccount *account.Account
}

func (a *ATM) ReadAccount(accMedium medium.Medium) error {
	acc, err := a.accountRepo.Account(accMedium.AccountID())
	if err != nil {
		return err
	}
	if err = acc.Authenticate(accMedium.Password()); err != nil {
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
