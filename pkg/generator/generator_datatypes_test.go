package generator

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerate_PrimitiveTypes(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []types.Token
		expected string
	}{
		{
			name: "String value",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Data"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "message"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "hello world"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"message":"hello world"}`,
		},
		{
			name: "Number value",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Data"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "score"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "100"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"score":100}`,
		},
		{
			name: "Boolean true",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Data"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "active"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "true"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"active":true}`,
		},
		{
			name: "Boolean false",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Data"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "deleted"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "false"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"deleted":false}`,
		},
		{
			name: "Null value",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Data"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "optional"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "null"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"optional":null}`,
		},
		{
			name: "Mixed data types",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Data"},
				{Type: types.PAREN_OPEN},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "Alice"},
				{Type: types.COMMA},
				{Type: types.KEY, Lexeme: "age"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "29"},
				{Type: types.COMMA},
				{Type: types.KEY, Lexeme: "verified"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "true"},
				{Type: types.COMMA},
				{Type: types.KEY, Lexeme: "middle_name"},
				{Type: types.EQUALS},
				{Type: types.VALUE, Lexeme: "null"},
				{Type: types.PAREN_CLOSE},
				{Type: types.EOF},
			},
			expected: `{"name":"Alice","age":29,"verified":true,"middle_name":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := Generate(tt.tokens)
			assert.JSONEq(t, tt.expected, string(result))
		})
	}
}
