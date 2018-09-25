package atm

import (
	"github.com/dyong0/atm/pkg/atm/account"
	"github.com/dyong0/atm/pkg/atm/account/medium"
	"github.com/dyong0/atm/pkg/atm/currency"
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
	err = a.accountRepo.VerifyAccount(acc, accMedium.Password())
	if err != nil {
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

func (a *ATM) Close() error {
	a.currentAccount = nil
	return nil
}

func NewATM() (*ATM, error) {
	repo, err := account.NewRepository()
	if err != nil {
		return nil, err
	}

	return &ATM{accountRepo: repo}, nil
}
