package token

import (
	"time"
)

func NewMemRepository(gen Generator) Repository {
	return &memRepository{
		tokens:    make(map[string]Token),
		generator: gen,
	}
}

type memRepository struct {
	Repository
	tokens    map[string]Token
	generator Generator
}

func (r *memRepository) Token(tokenId string) (Token, error) {
	t, ok := r.tokens[tokenId]
	if !ok {
		return Token{}, ErrTokenNotFound
	}

	return t, nil
}

func (r *memRepository) Generate() (Token, error) {
	newId, err := r.generator.Generate()
	if err != nil {
		return Token{}, err
	}

	t := Token{
		ID:            newId,
		Authorization: make(map[string]interface{}),
		ExpiresAt:     time.Now(),
	}

	return t, nil
}

func (r *memRepository) Save(token Token) error {
	r.tokens[token.ID] = token

	return nil
}
