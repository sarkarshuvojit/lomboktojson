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
	t.Run("Should scan simple empty doc with empty object", func(t *testing.T) {
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
	t.Run("Should scan Customer object with basic details", func(t *testing.T) {
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
	t.Run("Should scan customer with name: Tokens", func(t *testing.T) {
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
	t.Run("Should scan customer with name: TokenTypes", func(t *testing.T) {
		source := `Customer(name=Rajesh Kumar)`
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)
		scanner := lib.NewScanner(&sourceBuf)
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
	t.Run("Should scan customer with name, age, email: Tokens", func(t *testing.T) {
		source := `Customer(name=Rajesh Kumar,age=50,email=rajesh@kumar.com)`
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
	/*
		t.Skip("Should scan customer with name: TokenTypes", func(t *testing.T) {
			source := `Customer(name=Rajesh Kumar)`
			var sourceBuf bytes.Buffer
			sourceBuf.WriteString(source)
			scanner := lib.NewScanner(&sourceBuf)
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

		})*/
}
