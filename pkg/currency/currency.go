package currency

import (
	"errors"
	"sort"
)

type CurrencyKind int

var (
	CurrencyYen          = CurrencyKind(6)
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

func (a Amount) Add(depositAmount Amount) (Amount, error) {
	if a.kind != depositAmount.kind {
		return a, ErrDifferentCurrency
	}

	for k, v := range depositAmount.currencies {
		a.currencies[k] = a.currencies[k] + (depositAmount.currencies[k] * v)
	}

	return a, nil
}

func (a Amount) Subtract(depositAmount Amount) (Amount, error) {
	if a.kind != depositAmount.kind {
		return a, ErrDifferentCurrency
	}

	if a.Total() < depositAmount.Total() {
		return a, ErrMinusAmount
	}

	for k, v := range depositAmount.currencies {
		a.currencies[k] = a.currencies[k] - (depositAmount.currencies[k] * v)
	}

	return a, nil
}

func currenciesByKind(kind CurrencyKind) ([]uint32, error) {
	switch kind {
	case CurrencyYen:
		return YenCurrencies, nil
	default:
		return nil, ErrUnknownCurrency
	}
}

func New(kind CurrencyKind, total uint32) (Amount, error) {
	a := Amount{
		kind:       CurrencyYen,
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
	return New(CurrencyYen, total)
}
