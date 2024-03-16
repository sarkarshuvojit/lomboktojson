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
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   map[string]string
	line      int8
}

func (t Token) ToString() string {
	return fmt.Sprintf(
		"%d %s %s",
		t.tokenType, t.lexeme, t.literal,
	)
}
