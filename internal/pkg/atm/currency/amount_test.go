package currency

import "testing"

func TestAmount(t *testing.T) {
	_, err := NewAmount(CurrencyKind(-1), 0)
	if err != ErrUnknownCurrency {
		t.Error("Expect an error for unknown currency")
	}

	zeroAmount, _ := NewAmount(CurrencyKindYen, 0)
	if expect, got := uint32(0), zeroAmount.Total(); expect != got {
		t.Errorf("Expect zero amount to be %d, got %d", expect, got)
	}

	plusAmount, _ := NewAmount(CurrencyKindYen, 100)
	sumAmount, err := zeroAmount.Add(plusAmount)
	if err != nil {
		t.Error("Expect no error for adding correct amount")
	}
	if expect, got := uint32(100), sumAmount.Total(); expect != got {
		t.Errorf("Expect sum of %d, got %d", expect, got)
	}

	smallAmount, _ := NewAmount(CurrencyKindYen, 100)
	hugeAmount, _ := NewAmount(CurrencyKindYen, 100000)
	_, err = smallAmount.Subtract(hugeAmount)
	if err != ErrNegativeAmount {
		t.Errorf("Expect an error for subtracting %d from %d", hugeAmount.Total(), smallAmount.Total())
	}

	subtracted, err := hugeAmount.Subtract(smallAmount)
	if err != nil {
		t.Error("Expect no error for subtracting correct amount")
	}
	if expect, got := uint32(99900), subtracted.Total(); expect != got {
		t.Errorf("Expect subtraction result of %d, got %d", expect, got)
	}
}
