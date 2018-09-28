package generator

import (
	"strconv"
	"time"
)

func NewTimeBasedGenerator() *timeBasedGenerator {
	return &timeBasedGenerator{}
}

type timeBasedGenerator struct {
	Generator
}

func (g *timeBasedGenerator) Generate() (string, error) {
	return strconv.FormatInt(time.Now().UnixNano(), 10), nil
}
