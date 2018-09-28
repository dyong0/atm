package token

type Repository interface {
	Generate() (Token, error)
	Save(token Token) error
}
