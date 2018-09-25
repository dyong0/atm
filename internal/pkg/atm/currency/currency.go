package currency

import (
	"errors"
)

type CurrencyKind int

var (
	ErrInvalidAmount     = errors.New("Invalid amount of currencies")
	ErrUnknownCurrency   = errors.New("Unknown currency")
	ErrDifferentCurrency = errors.New("Different currency cannot be added or subtracted")
	ErrNegativeAmount    = errors.New("Minus amount")
)
