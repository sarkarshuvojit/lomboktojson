package lib

import (
	"bytes"
	"fmt"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

type SingleTokenToJson func (int, []types.Token) ([]byte, bool)

func getOptionallyQuotedValue(val string) string {
	return fmt.Sprintf("\"%s\"", val)
}

var tokenTypeGeneratorMapping map[types.TokenType]SingleTokenToJson = map[types.TokenType]SingleTokenToJson{
	types.EOF: func(i int, ts []types.Token) ([]byte, bool) {
		return []byte(``), false
	},
	types.CLASS_NAME: func(i int, ts []types.Token) ([]byte, bool) {
		return []byte(``), false
	},
	types.KEY: func(i int, ts []types.Token) ([]byte, bool) {
		t := ts[i]
		return []byte(getOptionallyQuotedValue(t.Lexeme)), true
	},
	types.VALUE: func(i int, ts []types.Token) ([]byte, bool) {
		t := ts[i]
		return []byte(getOptionallyQuotedValue(t.Lexeme)), true
	},
	types.EQUALS: func(i int, ts []types.Token) ([]byte, bool) {
		// t := ts[i]
		return []byte(`:`), true
	},
	types.COMMA: func(i int, ts []types.Token) ([]byte, bool) {
		// t := ts[i]
		return []byte(`,`), true
	},
	types.PAREN_OPEN: func(i int, t []types.Token) ([]byte, bool) {
		if i == 0 {
			return []byte(``), false
		}
		if t[i-1].Type == types.CLASS_NAME {
			return []byte(`{`), true
		}
		return []byte(``), false
	},
	types.PAREN_CLOSE: func(i int, t []types.Token) ([]byte, bool) {
		return []byte(`}`), true
	},
	types.ARRAY_OPEN: func(i int, t []types.Token) ([]byte, bool) {
		return []byte(`[`), true
	},
	types.ARRAY_CLOSE: func(i int, t []types.Token) ([]byte, bool) {
		return []byte(`]`), true
	},

}

func Generate(tokens []types.Token) ([]byte, error) {
	var buf bytes.Buffer
	for i := range tokens {
		if tknBytes, present := generateTokenAt(i, tokens); present {
			buf.Write(tknBytes)
		}
	}

	if buf.Len() == 0 {
		buf.WriteString("{}")
	}
	asBytes := buf.Bytes()
	return asBytes, nil
}

func generateTokenAt(i int, tokens []types.Token) ([]byte, bool) {
	token := tokens[i]
	if converterFn, ok := tokenTypeGeneratorMapping[token.Type]; ok {
		return converterFn(i, tokens)
	} else {
		return []byte(``), false
	}
	
}
