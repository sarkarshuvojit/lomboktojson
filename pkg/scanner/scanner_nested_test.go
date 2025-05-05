package scanner

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Test_ScanNestedCustomerWithArraysLombokStyle(t *testing.T) {
    source := `Customer(
        name=Rajesh Kumar,
        addresses=[
            Address(street=123 Main St, city=Bangalore, zip=560001),
            Address(street=456 Oak Ave, city=Mumbai, zip=400001)
        ],
        contacts=[
            Contact(type=email, value=rajesh@kumar.com),
            Contact(type=phone, numbers=[9876543210, 1234567890])
        ]
    )`

    t.Run("Nested Customer with Arrays (Lombok Style): Token Length & Lexemes", func(t *testing.T) {
        var sourceBuf bytes.Buffer
        sourceBuf.WriteString(source)
        scanner := NewScanner(&sourceBuf)
        tokens := scanner.Scan()

        expectedTokenLen := 71

        if len(tokens) != expectedTokenLen {
			fmt.Println(tokens)
            t.Errorf("Should return %d tokens including EOF, got %d", expectedTokenLen, len(tokens))
        }

        expectedTokens := []string{
            "Customer", "(",
            "name", "=", "Rajesh Kumar", ",",
            "addresses", "=", "[",
                "Address", "(", "street", "=", "123 Main St", ",", "city", "=", "Bangalore", ",", "zip", "=", "560001", ")", ",",
                "Address", "(", "street", "=", "456 Oak Ave", ",", "city", "=", "Mumbai", ",", "zip", "=", "400001", ")",
            "]", ",",
            "contacts", "=", "[",
                "Contact", "(", "type", "=", "email", ",", "value", "=", "rajesh@kumar.com", ")", ",",
                "Contact", "(", "type", "=", "phone", ",", "numbers", "=", "[", "9876543210", ",", "1234567890", "]",
                ")",
            "]",
            ")",
        }

        for i := 0; i < len(expectedTokens); i++ {
            if i >= len(tokens) {
                t.Errorf("Missing token at position %d, expected %s", i, expectedTokens[i])
                continue
            }
            if tokens[i].Lexeme != expectedTokens[i] {
                t.Errorf("Token #%d lexeme expected '%s' got '%s'", i+1, expectedTokens[i], tokens[i].Lexeme)
            }
        }
    })

    t.Run("Nested Customer with Arrays (Lombok Style): Token Types", func(t *testing.T) {
        var sourceBuf bytes.Buffer
        sourceBuf.WriteString(source)
        scanner := NewScanner(&sourceBuf)
        tokens := scanner.Scan()

        expectedTokenLen := 71

        if len(tokens) != expectedTokenLen {
			fmt.Println(tokens)
            t.Errorf("Should return %d tokens including EOF, got %d", expectedTokenLen, len(tokens))
        }

		expectedTokenTypes := []types.TokenType{
			// Customer(
			types.CLASS_NAME,    // "Customer"
			types.PAREN_OPEN,    // "("

			// name=Rajesh Kumar,
			types.KEY,           // "name"
			types.EQUALS,        // "="
			types.VALUE,         // "Rajesh Kumar"
			types.COMMA,         // ","

			// addresses=[
			types.KEY,           // "addresses"
			types.EQUALS,        // "="
			types.ARRAY_OPEN,    // "["

			// Address(street=123 Main St, city=Bangalore, zip=560001),
			types.CLASS_NAME,    // "Address"
			types.PAREN_OPEN,    // "("
			types.KEY,          // "street"
			types.EQUALS,       // "="
			types.VALUE,        // "123 Main St"
			types.COMMA,        // ","
			types.KEY,          // "city"
			types.EQUALS,       // "="
			types.VALUE,        // "Bangalore"
			types.COMMA,        // ","
			types.KEY,          // "zip"
			types.EQUALS,       // "="
			types.VALUE,        // "560001"
			types.PAREN_CLOSE,  // ")"
			types.COMMA,        // ","

			// Address(street=456 Oak Ave, city=Mumbai, zip=400001)
			types.CLASS_NAME,    // "Address"
			types.PAREN_OPEN,   // "("
			types.KEY,           // "street"
			types.EQUALS,        // "="
			types.VALUE,         // "456 Oak Ave"
			types.COMMA,        // ","
			types.KEY,          // "city"
			types.EQUALS,       // "="
			types.VALUE,        // "Mumbai"
			types.COMMA,        // ","
			types.KEY,         // "zip"
			types.EQUALS,      // "="
			types.VALUE,       // "400001"
			types.PAREN_CLOSE, // ")"

			// ],
			types.ARRAY_CLOSE,   // "]"
			types.COMMA,         // ","

			// contacts=[
			types.KEY,          // "contacts"
			types.EQUALS,       // "="
			types.ARRAY_OPEN,   // "["

			// Contact(type=email, value=rajesh@kumar.com),
			types.CLASS_NAME,   // "Contact"
			types.PAREN_OPEN,   // "("
			types.KEY,          // "type"
			types.EQUALS,       // "="
			types.VALUE,        // "email"
			types.COMMA,        // ","
			types.KEY,          // "value"
			types.EQUALS,       // "="
			types.VALUE,        // "rajesh@kumar.com"
			types.PAREN_CLOSE,  // ")"
			types.COMMA,        // ","

			// Contact(type=phone, numbers=[9876543210, 1234567890])
			types.CLASS_NAME,   // "Contact"
			types.PAREN_OPEN,   // "("
			types.KEY,          // "type"
			types.EQUALS,       // "="
			types.VALUE,        // "phone"
			types.COMMA,        // ","
			types.KEY,          // "numbers"
			types.EQUALS,       // "="
			types.ARRAY_OPEN,   // "["
			types.STRING_LITERAL,    // "9876543210"
			types.COMMA,    // ","
			types.STRING_LITERAL,   // "1234567890"
			types.ARRAY_CLOSE,  // "]"
			types.PAREN_CLOSE,  // ")"

			// ]
			types.ARRAY_CLOSE,  // "]"

			// )
			types.PAREN_CLOSE,  // ")"
			types.EOF,          // EOF
		}

        for i := 0; i < len(expectedTokenTypes); i++ {
            if i >= len(tokens) {
                t.Errorf("Missing token at position %d", i)
                continue
            }
            if tokens[i].Type != expectedTokenTypes[i] {
                t.Errorf("Token #%d type expected %v got %v", i+1, expectedTokenTypes[i], tokens[i].Type)
            }
        }
    })
}
