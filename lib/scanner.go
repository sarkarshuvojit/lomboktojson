package lib

import (
	"fmt"
	"io"
	"os"
	"unicode"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

type Scanner struct {
	sourceBytes []byte

	curline int
	start   int
	end     int

	parenOpen int

	// Literal Related flags
	literalStarted bool
	literalStart   int
	literalEnd     int

	tokens []types.Token
}

func NewScanner(source io.Reader) *Scanner {
	sourceBytes, err := io.ReadAll(source)
	if err != nil {
		fmt.Printf("Could not read %v: %v\n", source, err)
		os.Exit(1)
	}

	return &Scanner{
		sourceBytes: sourceBytes,
		curline:     1,
		start:       0,
		end:         0,

		parenOpen: 0,

		literalStarted: false,
		literalStart:   -1,
		literalEnd:     -1,
	}
}

func isAlpha(ch string) bool {
	return unicode.IsLetter([]rune(ch)[0])
}

func (s *Scanner) stringLiteralToToken(literal string) types.Token {
	if s.sourceBytes[s.literalEnd+1] == '(' {
		return types.NewToken(
			types.CLASS_NAME,
			string(literal),
			nil,
			s.curline,
		)
	}
	if s.sourceBytes[s.literalEnd+1] == '=' {
		return types.NewToken(
			types.KEY,
			string(literal),
			nil,
			s.curline,
		)
	}
	if s.sourceBytes[s.literalStart-1] == '=' {
		return types.NewToken(
			types.VALUE,
			string(literal),
			nil,
			s.curline,
		)
	}
	return types.NewToken(
		types.STRING_LITERAL,
		string(literal),
		nil,
		s.curline,
	)

}

func (s *Scanner) clearStringLiterals() {
	if s.literalStarted {
		literal := s.sourceBytes[s.literalStart : s.literalEnd+1]
		_token := types.NewToken(
			types.STRING_LITERAL,
			string(literal),
			nil,
			s.curline,
		)
		_token = s.stringLiteralToToken(string(literal))
		s.tokens = append(s.tokens, _token)
		s.literalStarted = false
		s.literalStart = -1
	}
}

func (s *Scanner) Scan() []types.Token {

	for chIdx := range s.sourceBytes {
		ch := string(s.sourceBytes[chIdx])
		switch ch {
		case "(":
			s.clearStringLiterals()
			_token := types.NewToken(
				types.PAREN_OPEN,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen++
			break
		case ")":
			s.clearStringLiterals()
			_token := types.NewToken(
				types.PAREN_CLOSE,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen--
			break
		case "=":
			s.clearStringLiterals()
			_token := types.NewToken(
				types.EQUALS,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen--
			break
		default:
			if isAlpha(ch) {
				if s.literalStarted {
					s.literalEnd = chIdx
				} else {
					s.literalStarted = true
					s.literalStart = chIdx
				}
			}
		}
		s.end++
	}

	s.tokens = append(s.tokens, types.NewToken(types.EOF, "", nil, s.curline))
	return s.tokens
}
