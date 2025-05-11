package generator

import (
	"bytes"
	"fmt"
	"regexp"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

type singleTokenToJson func(int, []types.Token) ([]byte, bool)

func isNumeric(s string) bool {
	re := regexp.MustCompile(`^\d+$`)
	return re.MatchString(s)
}

func isFloatingPoint(s string) bool {
	re := regexp.MustCompile(`^-?\d+(\.\d+)?$`)
	return re.MatchString(s)
}

func getOptionallyQuotedValue(val string) string {
	isNull := val == "null"
	isBool := val == "true" || val == "false"
	isNum := isNumeric(val)
	isFloat := isFloatingPoint(val)

	if isNull || isBool || isNum || isFloat {
		return fmt.Sprintf("%s", val)
	}
	return fmt.Sprintf("\"%s\"", val)
}

var tokenTypeGeneratorMapping map[types.TokenType]singleTokenToJson = map[types.TokenType]singleTokenToJson{
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

// Generate converts a sequence of parsed tokens into a JSON-formatted byte slice.
//
// It iterates through the provided tokens and builds a JSON object based on
// recognized key-value patterns or structured groupings.
//
// Returns a JSON byte slice on success, or an empty JSON object ("{}") if no valid tokens are found.
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
