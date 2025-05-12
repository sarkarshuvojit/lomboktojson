package beautifier_test

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/pkg/beautifier"
	"github.com/sarkarshuvojit/lomboktojson/types"
	"github.com/stretchr/testify/assert"
)

func TestBeautify(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []types.Token
		indent   int
		expected string
		skipTest bool
	}{
		{
			name: "Empty object for just EOF",
			tokens: []types.Token{
				{Type: types.EOF},
			},
			indent:   2,
			expected: "",
		},
		{
			name: "Simple object with one property",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "John"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.EOF},
			},
			indent: 2,
			expected: `Person(
  name=John
)`,
		},
		{
			name: "Object with multiple properties",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "John"},
				{Type: types.COMMA, Lexeme: ","},
				{Type: types.KEY, Lexeme: "age"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "30"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.EOF},
			},
			indent:   2,
			skipTest: true,
			expected: `Person(
  name=John,\n  age=30
)`,
		},
		{
			name: "Nested object",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"}, {Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "name"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "John"},
				{Type: types.COMMA, Lexeme: ","},
				{Type: types.KEY, Lexeme: "address"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.CLASS_NAME, Lexeme: "Address"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "street"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "123 Main St"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.EOF},
			},
			indent:   2,
			skipTest: true,
			expected: `Person(
  name=John,\n  address=Address(
    street=123 Main St
  )
)`,
		},
		{
			name: "Array of simple values",
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "scores"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.ARRAY_OPEN, Lexeme: "["},
				{Type: types.VALUE, Lexeme: "90"},
				{Type: types.COMMA, Lexeme: ","},
				{Type: types.VALUE, Lexeme: "85"},
				{Type: types.COMMA, Lexeme: ","},
				{Type: types.VALUE, Lexeme: "95"},
				{Type: types.ARRAY_CLOSE, Lexeme: "]"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.EOF},
			},
			indent:   2,
			skipTest: true,
			expected: `Person(
  scores=[
    90,\n    85,\n    95
  ]
)`,
		},
		{
			name:     "Array of objects",
			skipTest: true,
			tokens: []types.Token{
				{Type: types.CLASS_NAME, Lexeme: "Person"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "addresses"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.ARRAY_OPEN, Lexeme: "["},
				{Type: types.CLASS_NAME, Lexeme: "Address"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "street"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "123 Main St"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.COMMA, Lexeme: ","},
				{Type: types.CLASS_NAME, Lexeme: "Address"},
				{Type: types.PAREN_OPEN, Lexeme: "("},
				{Type: types.KEY, Lexeme: "street"},
				{Type: types.EQUALS, Lexeme: "="},
				{Type: types.VALUE, Lexeme: "456 Oak Ave"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.ARRAY_CLOSE, Lexeme: "]"},
				{Type: types.PAREN_CLOSE, Lexeme: ")"},
				{Type: types.EOF},
			},
			indent: 2,
			expected: `Person(
  addresses=[
    Address(
      street=123 Main St
    ),\n    Address(
      street=456 Oak Ave
    )
  ]
)`,
		},
	}

	for _, tt := range tests {
		if !tt.skipTest {
			t.Run(tt.name, func(t *testing.T) {
				result, _ := beautifier.Beautify(tt.tokens, tt.indent)
				assert.Equal(t, tt.expected, string(result))
			})
		}
	}
}
