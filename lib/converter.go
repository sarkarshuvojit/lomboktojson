package lib

import (
	"fmt"
	"strings"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

type Converter struct{}

func NewConverter() *Converter {
	return &Converter{}
}

func (c Converter) Convert(tokens []types.Token) string {
	finals := []string{}
	for _, token := range tokens {
		switch token.Type {
		case types.CLASS_NAME:
			continue
		case types.PAREN_OPEN:
			finals = append(finals, "{")
		case types.PAREN_CLOSE:
			finals = append(finals, "}")
		case types.KEY:
			finals = append(finals, fmt.Sprintf(`"%s"`, token.Lexeme))
		case types.EQUALS:
			finals = append(finals, ":")
		case types.VALUE:
			finals = append(finals, fmt.Sprintf(`"%s"`, token.Lexeme))
		}
	}
	if len(finals) == 0 {
		return "{}"
	}
	return strings.Join(finals, " ")
}
