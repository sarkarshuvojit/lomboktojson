package scanner

import (
	"bytes"
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Test_ScanSimple(t *testing.T) {
	t.Run("Should scan simple empty doc", func(t *testing.T) {
		source := ``
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)

		scanner := NewScanner(&sourceBuf)
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
		scanner := NewScanner(&sourceBuf)
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
		scanner := NewScanner(&sourceBuf)
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
		scanner := NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		// Customer | ParenOpen | Key | EQUALS | Value | PAREN_CLOSE | EOF
		// 0			1			2	3		4		5			6
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
		scanner := NewScanner(&sourceBuf)
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
		scanner := NewScanner(&sourceBuf)
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
	source := `Customer(name=Rajesh Kumar,age=50,email=rajesh@kumar.com)`
	t.Run("Customer(name, age, email): TokenLengh & Lexemes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := NewScanner(&sourceBuf)
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
		scanner := NewScanner(&sourceBuf)
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

func Test_ScanTheCustomerWithArrayField(t *testing.T) {
	source := `Customer(name=Rajesh Kumar,phones=[1234567890, 9876543210])`
	t.Run("Customer(name, phones[]): TokenLengh & Lexemes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 15
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
			return
		}

		expectedTokens := [15]string{
			"Customer", "(",
			"name", "=", "Rajesh Kumar", ",",
			"phones", "=", "[", "1234567890", ",", "9876543210", "]",
			")"}

		for i := 0; i < len(expectedTokens); i++ {
			expectedToken := expectedTokens[i]
			actualToken := tokens[i]
			if actualToken.Lexeme != expectedToken {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedToken, tokens[i].Lexeme)
			}
		}
	})
	t.Run("Customer(name, phones[]): TokenTypes", func(t *testing.T) {
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := NewScanner(&sourceBuf)
		tokens := scanner.Scan()

		expectedTokenLen := 15
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return %d token for EOF, got %d", expectedTokenLen, len(tokens))
			return
		}
		expectedTokenTypes := [15]types.TokenType{
			types.CLASS_NAME, types.PAREN_OPEN,
			types.KEY, types.EQUALS, types.VALUE, types.COMMA,
			types.KEY, types.EQUALS, types.ARRAY_OPEN, types.STRING_LITERAL, types.COMMA, types.STRING_LITERAL, types.ARRAY_CLOSE,
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
