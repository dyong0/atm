package token

import (
	"fmt"
)

type Token struct {
	fmt.Stringer
	Authorization interface{}
}
