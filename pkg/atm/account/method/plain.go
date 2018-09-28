package method

type Plain struct {
	Method

	id string
	pw string
}

func NewPlain(id, pw string) Plain {
	return Plain{id: id, pw: pw}
}

func (m Plain) AccountID() string {
	return m.id
}

func (m Plain) Password() string {
	return m.pw
}
