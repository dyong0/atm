package token

import (
	"time"
)

type Token struct {
	ID            string
	Authorization map[string]interface{}
	ExpiresAt     time.Time
}
