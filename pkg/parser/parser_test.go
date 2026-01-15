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

func TestParse_FailsOnTrailingCommaObject(t *testing.T) {
	tokens := []types.Token{
		{Type: types.CLASS_NAME, Lexeme: "Customer"},
		{Type: types.PAREN_OPEN, Lexeme: "("},
		{Type: types.KEY, Lexeme: "name"},
		{Type: types.EQUALS, Lexeme: "="},
		{Type: types.VALUE, Lexeme: "John"},
		{Type: types.COMMA, Lexeme: ","},
		{Type: types.PAREN_CLOSE, Lexeme: ")"},
		{Type: types.EOF},
	}

	_, err := Parse(tokens)

	assert.ErrorIs(t, err, ErrTrailingCommaObject)
}

func TestParse_FailsOnTrailingCommaArray(t *testing.T) {
	tokens := []types.Token{
		{Type: types.CLASS_NAME, Lexeme: "Basket"},
		{Type: types.PAREN_OPEN, Lexeme: "("},
		{Type: types.KEY, Lexeme: "items"},
		{Type: types.EQUALS, Lexeme: "="},
		{Type: types.ARRAY_OPEN, Lexeme: "["},
		{Type: types.VALUE, Lexeme: "apple"},
		{Type: types.COMMA, Lexeme: ","},
		{Type: types.ARRAY_CLOSE, Lexeme: "]"},
		{Type: types.PAREN_CLOSE, Lexeme: ")"},
		{Type: types.EOF},
	}

	_, err := Parse(tokens)

	assert.ErrorIs(t, err, ErrTrailingCommaArray)
}

func TestParse_FailsOnMismatchedArrayCloser(t *testing.T) {
	tokens := []types.Token{
		{Type: types.CLASS_NAME, Lexeme: "Customer"},
		{Type: types.PAREN_OPEN, Lexeme: "("},
		{Type: types.KEY, Lexeme: "scores"},
		{Type: types.EQUALS, Lexeme: "="},
		{Type: types.ARRAY_OPEN, Lexeme: "["},
		{Type: types.VALUE, Lexeme: "1"},
		{Type: types.COMMA, Lexeme: ","},
		{Type: types.VALUE, Lexeme: "2"},
		{Type: types.PAREN_CLOSE, Lexeme: ")"},
		{Type: types.EOF},
	}

	_, err := Parse(tokens)

	assert.ErrorIs(t, err, ErrMismatchedCloser)
}
