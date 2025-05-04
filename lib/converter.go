package lib

import (
	"bytes"
	"fmt"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

type TokenConverter func (types.Token) (byte, bool)

var tokenConverterMapping map[types.TokenType]TokenConverter = map[types.TokenType]TokenConverter{
	types.EOF: func(t types.Token) (byte, bool) {
		return 0, false
	},
	types.CLASS_NAME: func(t types.Token) (byte, bool) {
		return 0, false
	},
}

func Convert(tokens []types.Token) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, token := range tokens {
		if tkn, present := convertToken(token); present {
			buf.WriteByte(tkn)
		}
	}

	buf.WriteByte('}')
	asBytes := buf.Bytes()
	fmt.Println(string(asBytes))
	return asBytes, nil
}

func convertToken(token types.Token) (byte, bool) {
	if converterFn, ok := tokenConverterMapping[token.Type]; ok {
		return converterFn(token)
	} else {
		return 0, false
	}
}
