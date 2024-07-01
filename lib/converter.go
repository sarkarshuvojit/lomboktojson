package lib

import (
	"github.com/sarkarshuvojit/lomboktojson/types"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c Converter) Convert(tokens []types.Token) string {
	return "{}"
}
