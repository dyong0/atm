package currency

import (
	"errors"
)

type CurrencyKind int

var (
	ErrInvalidAmount     = errors.New("invalid amount of currencies")
	ErrUnknownCurrency   = errors.New("unknown currency")
	ErrDifferentCurrency = errors.New("unable to add or subtract using different currencies")
	ErrNegativeAmount    = errors.New("minus amount")
)
