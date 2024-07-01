package lib_test

import (
	"bytes"
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/lib"
	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Test_ScanSimple(t *testing.T) {
	t.Run("Should scan simple empty doc", func(t *testing.T) {
		source := ``
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)

		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 1
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return one token for EOF, got %d", len(tokens))
		}

		expectedTokenType := types.EOF
		if tokens[0].Type != expectedTokenType {
			t.Errorf("Should return EOF, got %s", tokens[0].Type)
		}
	})
	t.Run("Customer()", func(t *testing.T) {
		source := `Customer()`
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 4
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
		}

		expectedFirstToken := "Customer"
		if tokens[0].Lexeme != expectedFirstToken {
			t.Errorf("First token lexeme expected %s got %s", expectedFirstToken, tokens[0].Lexeme)
		}
	})
}

func Test_ScanTheCustomer(t *testing.T) {
	t.Run("Customer(name): Tokens", func(t *testing.T) {
		source := `Customer(name=Rajesh Kumar)`
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 7
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
		}

		expectedTokens := [7]string{"Customer", "(", "name", "=", "Rajesh Kumar", ")"}

		for i := 0; i < len(expectedTokens); i++ {
			expectedToken := expectedTokens[i]
			actualToken := tokens[i]
			if actualToken.Lexeme != expectedToken {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedToken, tokens[i].Lexeme)
			}
		}
	})
	t.Run("Customer(name): TokenTypes", func(t *testing.T) {
		source := `Customer(name=Rajesh Kumar)`
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 7
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
		}

		expectedTokenTypes := [7]types.TokenType{types.CLASS_NAME, types.PAREN_OPEN, types.KEY, types.EQUALS, types.VALUE, types.PAREN_CLOSE, types.EOF}

		for i := 0; i < len(expectedTokenTypes); i++ {
			expectedTokenType := expectedTokenTypes[i]
			actualToken := tokens[i]
			if actualToken.Type != expectedTokenType {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedTokenType, tokens[i].Type)
			}
		}

	})
}

func Test_ScanTheCustomerWithMultipleFieldsFlat(t *testing.T) {
	source := `Customer(name=Rajesh Kumar,age=50,email=rajesh@kumar.com)`
	t.Run("Customer(name, age, email): TokenLengh & Lexemes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 15
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
		}

		expectedTokens := [15]string{
			"Customer", "(",
			"name", "=", "Rajesh Kumar", ",",
			"age", "=", "50", ",",
			"email", "=", "rajesh@kumar.com",
			")"}

		for i := 0; i < len(expectedTokens); i++ {
			expectedToken := expectedTokens[i]
			actualToken := tokens[i]
			if actualToken.Lexeme != expectedToken {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedToken, tokens[i].Lexeme)
			}
		}
	})
	t.Run("Customer(name, age, email): TokenTypes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenTypes := [15]types.TokenType{
			types.CLASS_NAME, types.PAREN_OPEN,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.VALUE,
			types.PAREN_CLOSE,
			types.EOF}

		for i := 0; i < len(expectedTokenTypes); i++ {
			expectedTokenType := expectedTokenTypes[i]
			actualToken := tokens[i]
			if actualToken.Type != expectedTokenType {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedTokenType, tokens[i].Type)
			}
		}

	})
}

func Test_ScanTheCustomerWithMultipleFieldsNested(t *testing.T) {
	source := `Customer(name=Michael Scott, age=40, email=michael@dundermifflin.com, address=Address(billingAddress=Some random address, shippingAddress=Some random address again))`
	t.Run("Customer(name, age, email, Address(billingAddress,shippingAddress)): TokenLengh & Lexemes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 28
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
		}

		expectedTokens := [28]string{
			"Customer", "(",
			"name", "=", "Michael Scott", ",",
			"age", "=", "40", ",",
			"email", "=", "michael@dundermifflin.com", ",",
			"address", "=",
			"Address", "(",
			"billingAddress", "=", "Some random address", ",",
			"shippingAddress", "=", "Some random address again",
			")",
			")"}

		for i := 0; i < len(expectedTokens); i++ {
			expectedToken := expectedTokens[i]
			actualToken := tokens[i]
			if actualToken.Lexeme != expectedToken {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedToken, tokens[i].Lexeme)
			}
		}
	})
	t.Run("Customer(name, age, email, Address(billingAddress,shippingAddress)): TokenTypes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenTypes := [28]types.TokenType{
			types.CLASS_NAME, types.PAREN_OPEN,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.CLASS_NAME, types.PAREN_OPEN,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.VALUE,
			types.PAREN_CLOSE, types.PAREN_CLOSE,
			types.EOF}

		for i := 0; i < len(expectedTokenTypes); i++ {
			expectedTokenType := expectedTokenTypes[i]
			actualToken := tokens[i]
			if actualToken.Type != expectedTokenType {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedTokenType, tokens[i].Type)
			}
		}

	})
}
