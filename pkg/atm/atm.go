package atm

import "github.com/dyong0/atm/pkg/account"

type ATM struct {
	accountService account.Service
}

func (a *ATM) ReadAccount(id string, pw string) error {
	err := a.accountService.ReadAccount(id, pw)
	if err != nil {
		return err
	}

	return nil
}

func (a *ATM) Deposit() {

}

func (a *ATM) Withdraw() {

}

func New() *ATM {
	return &ATM{
		accountService: account.NewService(),
	}
}
