package types

import "fmt"

type TokenType uint

const (
	PAREN_OPEN     TokenType = iota
	PAREN_CLOSE    TokenType = iota
	EQUALS         TokenType = iota
	COMMA          TokenType = iota
	EOF            TokenType = iota
	STRING_LITERAL TokenType = iota
	NUM_LITERAL    TokenType = iota
	CLASS_NAME     TokenType = iota
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal map[string]string
	Line    int
}

func (t Token) ToString() string {
	return fmt.Sprintf(
		"%d %s %s",
		t.Type, t.Lexeme, t.Literal,
	)
}

func NewToken(
	tokenType TokenType,
	lexeme string,
	literal map[string]string,
	line int,
) Token {
	return Token{
		Type:    tokenType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}
}
