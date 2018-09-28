package method

type AccountBook struct {
	Method
	accountID string
	password  string
}

func (b AccountBook) AccountID() string {
	return b.accountID
}

func (b AccountBook) Password() string {
	return b.password
}
