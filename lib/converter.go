package lib

import (
	"bytes"
	"fmt"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

type TokenConverter func (types.Token) ([]byte, bool)

func getOptionallyQuotedValue(val string) string {
	return fmt.Sprintf("\"%s\"", val)
}

var tokenConverterMapping map[types.TokenType]TokenConverter = map[types.TokenType]TokenConverter{
	types.EOF: func(t types.Token) ([]byte, bool) {
		return []byte(``), false
	},
	types.CLASS_NAME: func(t types.Token) ([]byte, bool) {
		return []byte(``), false
	},
	types.KEY: func(t types.Token) ([]byte, bool) {
		return []byte(getOptionallyQuotedValue(t.Lexeme)), true
	},
	types.VALUE: func(t types.Token) ([]byte, bool) {
		return []byte(getOptionallyQuotedValue(t.Lexeme)), true
	},
	types.EQUALS: func(t types.Token) ([]byte, bool) {
		return []byte(`:`), true
	},
	types.COMMA: func(t types.Token) ([]byte, bool) {
		return []byte(`,`), true
	},

}

func Convert(tokens []types.Token) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, token := range tokens {
		if tknBytes, present := convertToken(token); present {
			buf.Write(tknBytes)
		}
	}

	buf.WriteByte('}')
	asBytes := buf.Bytes()
	fmt.Println(string(asBytes))
	return asBytes, nil
}

func convertToken(token types.Token) ([]byte, bool) {
	if converterFn, ok := tokenConverterMapping[token.Type]; ok {
		return converterFn(token)
	} else {
		return []byte(``), false
	}
}
