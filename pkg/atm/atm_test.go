package atm

import (
	"testing"

	"github.com/dyong0/atm/pkg/atm/account"
	"github.com/dyong0/atm/pkg/atm/account/medium"
	"github.com/dyong0/atm/pkg/atm/currency"
)

func TestATM(t *testing.T) {
	atm, _ := NewATM()

	atm.accountRepo = &mockAccRepo{
		AccountFunc: func(id string) (*account.Account, error) {
			return account.NewAccount(currency.CurrencyYen), nil
		},
		VerifyAccountFunc: func(account *account.Account, pw string) error {
			return nil
		},
	}

	atm.ReadAccount(card{id: "awesome", pw: "secret"})

	depositAmount, _ := currency.Yen(6000)
	atm.Deposit(depositAmount)
	if expect, got := uint32(6000), atm.Balance(); expect != got {
		t.Errorf("Expect deposit amount of %d, got %d", expect, got)
	}

	withdrawAmount, _ := currency.Yen(651)
	withdrawn, err := atm.Withdraw(withdrawAmount)
	if expect, got := uint32(651), withdrawn.Total(); expect != got {
		t.Errorf("Expect withdrawn %d, got %d", expect, got)
	}
	if err != nil {
		t.Error("Expect no error while withdrawing")
	}
	if expect, got := uint32(6000-651), atm.Balance(); expect != got {
		t.Errorf("Expect balance of %d, got %d", expect, got)
	}

	err = atm.Close()
	if err != nil {
		t.Error("Expect to close ATM successfully")
	}
}

type mockAccRepo struct {
	AccountFunc       func(id string) (*account.Account, error)
	VerifyAccountFunc func(account *account.Account, pw string) error
}

func (r *mockAccRepo) Account(id string) (*account.Account, error) {
	return r.AccountFunc(id)
}
func (r *mockAccRepo) VerifyAccount(account *account.Account, pw string) error {
	return r.VerifyAccountFunc(account, pw)
}

type card struct {
	medium.Medium
	id string
	pw string
}

func (c card) AccountID() string {
	return c.id
}
func (c card) Password() string {
	return c.pw
}
