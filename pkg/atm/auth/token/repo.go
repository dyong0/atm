package token

import "errors"

var (
	ErrTokenNotFound = errors.New("token not found")
)

type Repository interface {
	Token(tokenId string) (Token, error)
	Generate() (Token, error)
	Save(token Token) error
}
