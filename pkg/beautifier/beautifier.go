package beautifier

import (
	"bytes"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

func tabs(indent int, curtab int) string {
	s := ""
	for i := 0; i < indent*curtab; i++ {
		s += " "
	}
	return s
}

func shouldIndent(token types.Token) bool {
	return !(token.Type == types.EQUALS || token.Type == types.VALUE ||
		token.Type == types.ARRAY_OPEN || token.Type == types.ARRAY_CLOSE ||
		token.Type == types.PAREN_OPEN || // token.Type == types.PAREN_CLOSE ||
		token.Type == types.COMMA || token.Type == types.CLASS_NAME)
}

func shouldPrintNewLineAfter(tokens []types.Token, i int) bool {
	token := tokens[i]

	if i == len(tokens)-2 {
		// Last Token
		return false
	}

	if token.Type == types.VALUE && tokens[i+1].Type != types.COMMA {
		// Last key of object before paren end or array end
		return true
	}

	return token.Type == types.COMMA ||
		token.Type == types.PAREN_OPEN ||
		token.Type == types.PAREN_CLOSE ||
		token.Type == types.ARRAY_OPEN ||
		token.Type == types.ARRAY_CLOSE
}

func shouldPrintNewLineBefore(token types.Token) bool {
	return false
}

func shouldIncreaseIndent(token types.Token) bool {
	return token.Type == types.PAREN_OPEN || token.Type == types.ARRAY_OPEN
}
func shouldDecreaseIndent(token types.Token) bool {
	return token.Type == types.PAREN_CLOSE || token.Type == types.ARRAY_CLOSE
}

func Beautify(tokens []types.Token, indent int) (asBytes []byte, err error) {
	var buf bytes.Buffer
	curTab := 0
	for i := range tokens {
		token := tokens[i]
		if shouldDecreaseIndent(token) {
			curTab--
		}

		if shouldIndent(token) {
			buf.WriteString(tabs(indent, curTab))
		}

		if shouldPrintNewLineBefore(token) {
			buf.WriteString("\n")
		}

		// Write the actual token's lexeme
		buf.WriteString(token.Lexeme)

		if shouldIncreaseIndent(token) {
			curTab++
		}

		if shouldPrintNewLineAfter(tokens, i) {
			buf.WriteString("\n")
		}

	}
	asBytes = buf.Bytes()
	return asBytes, nil
}
