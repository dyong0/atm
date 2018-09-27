package method

type Card struct {
	Method
	cardNo    string
	accountID string
	password  string
}

func NewCard(cardNo, password string) Card {
	return Card{
		cardNo:    cardNo,
		accountID: cardNo,
		password:  password,
	}
}

func (c Card) AccountID() string {
	return c.accountID
}
func (c Card) Password() string {
	return c.password
}
