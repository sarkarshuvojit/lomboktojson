package beautifier

import (
	"bytes"
	"strings"

	"github.com/sarkarshuvojit/lomboktojson/pkg/scanner"
	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Beautify(tokens []types.Token, indent int) (asBytes []byte, err error) {
	if indent <= 0 {
		indent = 1
	}

	var buf bytes.Buffer
	tabs := 0
	needIndent := true

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token.Type == types.EOF {
			break
		}

		switch token.Type {
		case types.PAREN_CLOSE, types.ARRAY_CLOSE:
			if tabs > 0 {
				tabs--
			}
			if !needIndent {
				buf.WriteString("\n")
			}
			needIndent = true
		}

		if needIndent {
			buf.WriteString(strings.Repeat(" ", indent*tabs))
			needIndent = false
		}

		buf.WriteString(token.Lexeme)

		switch token.Type {
		case types.CLASS_NAME:
			if i+1 < len(tokens) && tokens[i+1].Type == types.PAREN_OPEN {
				// handled alongside the following paren
			}
		case types.PAREN_OPEN, types.ARRAY_OPEN:
			tabs++
			buf.WriteString("\n")
			needIndent = true
		case types.COMMA:
			buf.WriteString("\n")
			needIndent = true
		case types.PAREN_CLOSE, types.ARRAY_CLOSE:
			if i+1 < len(tokens) && tokens[i+1].Type == types.COMMA {
				// comma will handle newline
			} else if i+1 < len(tokens) && (tokens[i+1].Type == types.PAREN_CLOSE || tokens[i+1].Type == types.ARRAY_CLOSE) {
				buf.WriteString("\n")
				needIndent = true
			} else if i+1 < len(tokens) && tokens[i+1].Type != types.EOF {
				buf.WriteString("\n")
				needIndent = true
			}
		}
	}

	result := strings.TrimRight(buf.String(), "\n")
	return []byte(result), nil
}

// BeautifySource scans the input string and returns the formatted output.
func BeautifySource(source string, indent int) (string, error) {
	sc := scanner.NewScanner(strings.NewReader(source))
	tokens := sc.Scan()
	formatted, err := Beautify(tokens, indent)
	if err != nil {
		return "", err
	}
	return string(formatted), nil
}
