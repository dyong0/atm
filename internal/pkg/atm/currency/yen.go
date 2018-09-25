package currency

const (
	CurrencyKindYen        = CurrencyKind(6)
	Yen1            uint32 = 1
	Yen5            uint32 = 5
	Yen10           uint32 = 10
	Yen50           uint32 = 50
	Yen100          uint32 = 100
	Yen1000         uint32 = 1000
	Yen5000         uint32 = 5000
	Yen10000        uint32 = 10000
)

var YenCurrencies = []uint32{
	Yen1,
	Yen5,
	Yen10,
	Yen50,
	Yen100,
	Yen1000,
	Yen5000,
	Yen10000,
}

func Yen(total uint32) (Amount, error) {
	return NewAmount(CurrencyKindYen, total)
}
