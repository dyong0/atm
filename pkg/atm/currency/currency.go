package currency

import (
	"errors"
	"sort"
)

type CurrencyKind int

var (
	CurrencyKindYen      = CurrencyKind(6)
	ErrInvalidAmount     = errors.New("Invalid amount of currencies")
	ErrUnknownCurrency   = errors.New("Unknown currency")
	ErrDifferentCurrency = errors.New("Different currency cannot be added or subtracted")
	ErrMinusAmount       = errors.New("Minus amount")
)

type Amount struct {
	kind       CurrencyKind
	currencies map[uint32]uint32
}

func (a Amount) Total() uint32 {
	var total uint32

	for k, v := range a.currencies {
		total = total + (k * v)
	}

	return total
}

func (a Amount) Add(amount Amount) (Amount, error) {
	if a.kind != amount.kind {
		return a, ErrDifferentCurrency
	}

	newAmount, err := New(a.kind, a.Total()+amount.Total())
	if err != nil {
		return a, ErrInvalidAmount
	}

	return newAmount, nil
}

func (a Amount) Subtract(amount Amount) (Amount, error) {
	if a.kind != amount.kind {
		return a, ErrDifferentCurrency
	}

	if a.Total() < amount.Total() {
		return a, ErrMinusAmount
	}

	newAmount, err := New(a.kind, a.Total()-amount.Total())
	if err != nil {
		return a, ErrInvalidAmount
	}
	return newAmount, nil
}

func currenciesByKind(kind CurrencyKind) ([]uint32, error) {
	switch kind {
	case CurrencyKindYen:
		return YenCurrencies, nil
	default:
		return nil, ErrUnknownCurrency
	}
}

func New(kind CurrencyKind, total uint32) (Amount, error) {
	a := Amount{
		kind:       CurrencyKindYen,
		currencies: make(map[uint32]uint32),
	}

	currencies, err := currenciesByKind(kind)
	if err != nil {
		return a, err
	}

	sort.Slice(currencies, func(i, j int) bool { return currencies[i] > currencies[j] })

	for _, c := range currencies {
		for total >= c {
			total = total - c
			a.currencies[c]++
		}
	}

	if total < 0 {
		return a, ErrInvalidAmount
	}

	return a, nil
}

func Yen(total uint32) (Amount, error) {
	return New(CurrencyKindYen, total)
}
