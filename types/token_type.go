package types

import "fmt"

type TokenType string

const (
	PAREN_OPEN     TokenType = "PAREN_OPEN"
	PAREN_CLOSE    TokenType = "PAREN_CLOSE"
	EQUALS         TokenType = "EQUALS"
	COMMA          TokenType = "COMMA"
	EOF            TokenType = "EOF"
	STRING_LITERAL TokenType = "STRING_LITERAL"
	NUM_LITERAL    TokenType = "NUM_LITERAL"
	CLASS_NAME     TokenType = "CLASS_NAME"
	KEY            TokenType = "KEY"
	VALUE          TokenType = "VALUE"
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal map[string]string
	Line    int
}

func (t Token) ToString() string {
	return fmt.Sprintf(
		"%s %s %s",
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
