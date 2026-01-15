package parser

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/types"
	"github.com/stretchr/testify/assert"
)

func TestParse_FailsOnTrailingTokens(t *testing.T) {
	tokens := []types.Token{
		{Type: types.CLASS_NAME, Lexeme: "Customer"},
		{Type: types.PAREN_OPEN, Lexeme: "("},
		{Type: types.KEY, Lexeme: "name"},
		{Type: types.EQUALS, Lexeme: "="},
		{Type: types.VALUE, Lexeme: "John"},
		{Type: types.PAREN_CLOSE, Lexeme: ")"},
		{Type: types.STRING_LITERAL, Lexeme: "extra"},
		{Type: types.EOF},
	}

	_, err := Parse(tokens)

	assert.Error(t, err)
}
