package account

import (
	"testing"

	"github.com/dyong0/atm/pkg/atm/currency"
)

func TestAccount(t *testing.T) {
	acc := NewAccount(currency.CurrencyKindYen)

	amount, err := currency.Yen(1110)

	if err != nil {
		t.Errorf("Failed to create money %v", err)
	}

	err = acc.Deposit(amount)

	if err != nil {
		t.Errorf("Failed to deposit: %v", err)
	}

	if expect, got := uint32(1110), acc.Balance(); expect != got {
		t.Errorf("Expect %d, got %d", expect, got)
	}

	withdrawAmount, _ := currency.Yen(100000)
	_, err = acc.Withdraw(withdrawAmount)

	if err == nil {
		t.Errorf("Withdrawing more than the balance must fail")
	}

	withdrawAmount, _ = currency.Yen(50)
	withdrawn, err := acc.Withdraw(withdrawAmount)

	if err != nil {
		t.Errorf("Withdrawing appropriate amount must not fail: %v", err)
	}

	if expect, got := uint32(50), withdrawn.Total(); expect != got {
		t.Errorf("Expect to witdraw %d, got %d", expect, got)
	}

	if expect, got := uint32(1060), acc.Balance(); expect != got {
		t.Errorf("Expect remaining %d, got %d", expect, got)
	}
}
