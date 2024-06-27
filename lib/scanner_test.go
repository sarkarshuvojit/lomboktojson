package lib_test

import (
	"bytes"
	"fmt"
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
			t.Errorf("Should return EOF, got %d", tokens[0].Type)
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
	t.Run("Should scan customer with name", func(t *testing.T) {
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

		expectedTokens := [7]string{"Customer", "(", "name", "=", "Rajesh Kumar", ")"}

		for i := 0; i < len(expectedTokens); i++ {
			expectedToken := expectedTokens[i]
			actualToken := tokens[i]
			if actualToken.Lexeme != expectedToken {
				t.Errorf("Token #%d lexeme expected %s got %s", i+1, expectedToken, tokens[0].Lexeme)
			}
		}

		fmt.Println(tokens)
	})
}
