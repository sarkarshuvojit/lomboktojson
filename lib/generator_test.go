package lib_test

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/lib"
	"github.com/sarkarshuvojit/lomboktojson/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []types.Token
		expected string
	}{
		{
			name: "Empty object for just EOF",
			tokens: []types.Token{
				{Type: types.EOF},
			},
			expected: "{}",
		},
		{
			name: "Simple object with one property",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "John"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"name":"John"}`,
		},
		{
			name: "Object with multiple properties",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "John"},
				{Type: types.COMMA},
				{Type: types.KEY, Lexeme: "age"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "30"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"name":"John","age":"30"}`,
		},
		{
			name: "Nested object",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"}, {Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "John"},
				{Type: types.COMMA},
				{Type: types.KEY, Lexeme: "address"},
				{Type: types.EQUALS},
				{Type: types.CLASS_NAME, Lexeme: "Address"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "street"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "123 Main St"},
				{Type: types.PAREN_CLOSE},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"name":"John","address":{"street":"123 Main St"}}`,
		},
		{
			name: "Array of simple values",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "scores"},
				{Type: types.EQUALS},
				{Type: types.ARRAY_OPEN},
				{Type: types.VALUE, Lexeme: "90"},
				{Type: types.COMMA},
				{Type: types.VALUE, Lexeme: "85"},
				{Type: types.COMMA},
				{Type: types.VALUE, Lexeme: "95"},
				{Type: types.ARRAY_CLOSE},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"scores":["90","85","95"]}`,
		},
		{
			name: "Array of objects",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "addresses"},
				{Type: types.EQUALS},
				{Type: types.ARRAY_OPEN},
				{Type: types.CLASS_NAME, Lexeme: "Address"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "street"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "123 Main St"},
				{Type: types.PAREN_CLOSE},
				{Type: types.COMMA},
				{Type: types.CLASS_NAME, Lexeme: "Address"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "street"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "456 Oak Ave"},
				{Type: types.PAREN_CLOSE},
				{Type: types.ARRAY_CLOSE},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"addresses":[{"street":"123 Main St"},{"street":"456 Oak Ave"}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := lib.Generate(tt.tokens)
			assert.JSONEq(t, tt.expected, string(result))
		})
	}
}
