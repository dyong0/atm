package atm

import (
	"testing"

	"github.com/dyong0/atm/internal/pkg/atm/account"
	"github.com/dyong0/atm/internal/pkg/atm/account/method"
	"github.com/dyong0/atm/internal/pkg/atm/currency"
)

func TestATM(t *testing.T) {
	mockRepo := &mockAccRepo{
		AccountFunc: func(id string) (*account.Account, error) {
			return account.NewAccount(currency.CurrencyKindYen), nil
		},
	}
	atm, _ := NewATM(mockRepo)

	atm.ReadAccount(mockAccMethod{})

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
	account.Repository
	AccountFunc func(id string) (*account.Account, error)
}

func (r *mockAccRepo) Account(id string) (*account.Account, error) {
	return r.AccountFunc(id)
}

type mockAccMethod struct {
	method.Method
}

func (m mockAccMethod) AccountID() string {
	return ""
}

func (m mockAccMethod) Password() string {
	return ""
}
