package token

type Repository interface {
	Token(tokenId string) (Token, error)
	Generate() (Token, error)
	Save(token Token) error
}
