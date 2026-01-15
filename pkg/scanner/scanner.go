package scanner

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

	containerStack []types.TokenType
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

		containerStack: []types.TokenType{},
	}
}

func isAlpha(ch string) bool {
	return unicode.IsLetter([]rune(ch)[0])
}

func isNum(ch string) bool {
	// TODO: Optimisation: Search whole string
	return unicode.IsDigit([]rune(ch)[0])
}

func isLiteral(ch string) bool {
	// Allow '.' to stay within a literal so we keep truncated floats intact (e.g., 999.99...)
	// and let later numeric validation decide whether to quote or not.
	return isAlpha(ch) || isNum(ch) || ch == "."
}

func (s *Scanner) lastToken() *types.Token {
	if len(s.tokens) == 0 {
		return nil
	}
	return &s.tokens[len(s.tokens)-1]
}

func (s *Scanner) ensureValuePresent() error {
	last := s.lastToken()
	if last != nil && last.Type == types.EQUALS {
		return ErrValueExpected
	}
	return nil
}

func (s *Scanner) pushContainer(t types.TokenType) {
	s.containerStack = append(s.containerStack, t)
}

func (s *Scanner) popContainer() {
	if len(s.containerStack) == 0 {
		return
	}
	s.containerStack = s.containerStack[:len(s.containerStack)-1]
}

func (s *Scanner) currentContainer() types.TokenType {
	if len(s.containerStack) == 0 {
		return ""
	}
	return s.containerStack[len(s.containerStack)-1]
}

func (s *Scanner) stringLiteralToToken(literal string) types.Token {
	if s.literalEnd+1 < len(s.sourceBytes) && s.sourceBytes[s.literalEnd+1] == '(' {
		return types.NewToken(
			types.CLASS_NAME,
			string(literal),
			nil,
			s.curline,
		)
	}
	if s.literalEnd+1 < len(s.sourceBytes) && s.sourceBytes[s.literalEnd+1] == '=' {
		return types.NewToken(
			types.KEY,
			string(literal),
			nil,
			s.curline,
		)
	}
	if s.literalStart-1 >= 0 && s.sourceBytes[s.literalStart-1] == '=' {
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

func (s *Scanner) clearStringLiterals() error {
	if s.literalStarted {
		if s.literalEnd < s.literalStart {
			s.literalEnd = s.literalStart
		}
		prevToken := s.lastToken()
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

		// If we're inside an object and just saw a bare literal where a key should be, error out.
		if _token.Type == types.STRING_LITERAL && s.currentContainer() == types.PAREN_OPEN {
			if prevToken == nil || prevToken.Type == types.PAREN_OPEN || prevToken.Type == types.COMMA {
				return ErrKeyExpected
			}
		}
	}
	return nil
}

// Scan processes the input source and returns a slice of tokens.
//
// It walks through the byte stream, identifies literals, delimiters,
// and structural characters, and builds a tokenized representation
// of the Lombok-formatted string.
func (s *Scanner) Scan() ([]types.Token, error) {

	for chIdx := range s.sourceBytes {
		ch := string(s.sourceBytes[chIdx])
		switch ch {
		case "(":
			if err := s.clearStringLiterals(); err != nil {
				return nil, err
			}
			_token := types.NewToken(
				types.PAREN_OPEN,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen++
			s.pushContainer(types.PAREN_OPEN)
			break
		case ")":
			if err := s.clearStringLiterals(); err != nil {
				return nil, err
			}
			_token := types.NewToken(
				types.PAREN_CLOSE,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen--
			s.popContainer()
			break
		case "=":
			if err := s.clearStringLiterals(); err != nil {
				return nil, err
			}
			last := s.lastToken()
			if last == nil || last.Type != types.KEY {
				return nil, ErrKeyExpected
			}
			_token := types.NewToken(
				types.EQUALS,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			break
		case ",":
			if err := s.clearStringLiterals(); err != nil {
				return nil, err
			}
			if err := s.ensureValuePresent(); err != nil {
				return nil, err
			}
			_token := types.NewToken(
				types.COMMA,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			break
		case "[":
			s.clearStringLiterals()
			_token := types.NewToken(
				types.ARRAY_OPEN,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen++
			s.pushContainer(types.ARRAY_OPEN)
			break
		case "]":
			if err := s.clearStringLiterals(); err != nil {
				return nil, err
			}
			if err := s.ensureValuePresent(); err != nil {
				return nil, err
			}
			_token := types.NewToken(
				types.ARRAY_CLOSE,
				ch,
				nil,
				s.curline,
			)
			s.tokens = append(s.tokens, _token)
			s.parenOpen--
			s.popContainer()
			break
		default:
			if isLiteral(ch) {
				if s.literalStarted {
					s.literalEnd = chIdx
				} else {
					s.literalStarted = true
					s.literalStart = chIdx
					s.literalEnd = chIdx
				}
			}
		}
		s.end++
	}

	if err := s.clearStringLiterals(); err != nil {
		return nil, err
	}
	if err := s.ensureValuePresent(); err != nil {
		return nil, err
	}

	s.tokens = append(s.tokens, types.NewToken(types.EOF, "", nil, s.curline))
	return s.tokens, nil
}
