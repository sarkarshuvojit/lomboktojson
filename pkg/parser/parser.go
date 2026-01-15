package parser

import (
	"fmt"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

// Parser transforms scanned tokens into an AST representation.
type Parser struct {
	tokens []types.Token
	pos    int
}

// Parse turns the provided tokens into an AST root node.
func Parse(tokens []types.Token) (types.Node, error) {
	p := &Parser{tokens: tokens}
	if p.isAtEnd() {
		return nil, nil
	}

	node, err := p.parseNode()
	if err != nil {
		return nil, err
	}
	if !p.isAtEnd() {
		return nil, fmt.Errorf("%w: %s at line %d", ErrTrailingTokens, p.peek().Type, p.peek().Line)
	}
	return node, nil
}

func (p *Parser) parseNode() (types.Node, error) {
	if p.isAtEnd() {
		return nil, ErrUnexpectedEOF
	}

	switch p.peek().Type {
	case types.CLASS_NAME:
		return p.parseObject()
	case types.ARRAY_OPEN:
		return p.parseArray()
	case types.VALUE, types.STRING_LITERAL:
		tok := p.advance()
		return &types.ValueNode{Value: tok.Lexeme}, nil
	default:
		return nil, fmt.Errorf("%w: %s at line %d", ErrUnexpectedToken, p.peek().Type, p.peek().Line)
	}
}

func (p *Parser) parseObject() (types.Node, error) {
	classTok, err := p.consume(types.CLASS_NAME, "expected class name")
	if err != nil {
		return nil, err
	}
	if _, err = p.consume(types.PAREN_OPEN, "expected '(' after class name"); err != nil {
		return nil, err
	}

	fields := []types.ObjectField{}
	for !p.check(types.PAREN_CLOSE) {
		keyTok, err := p.consume(types.KEY, "expected key")
		if err != nil {
			return nil, err
		}
		if _, err = p.consume(types.EQUALS, "expected '=' after key"); err != nil {
			return nil, err
		}
		val, err := p.parseNode()
		if err != nil {
			return nil, err
		}
		fields = append(fields, types.ObjectField{
			Key:   keyTok.Lexeme,
			Value: val,
		})

		if p.match(types.COMMA) {
			if p.check(types.PAREN_CLOSE) {
				return nil, ErrTrailingCommaObject
			}
			continue
		}
		if p.check(types.PAREN_CLOSE) {
			break
		}
		if p.isAtEnd() {
			return nil, ErrUnexpectedEOF
		}
		return nil, fmt.Errorf("%w: %s at line %d", ErrUnexpectedToken, p.peek().Type, p.peek().Line)
	}

	if _, err = p.consume(types.PAREN_CLOSE, "expected ')' after object body"); err != nil {
		return nil, err
	}

	return &types.ObjectNode{
		ClassName: classTok.Lexeme,
		Fields:    fields,
	}, nil
}

func (p *Parser) parseArray() (types.Node, error) {
	if _, err := p.consume(types.ARRAY_OPEN, "expected '[' to start array"); err != nil {
		return nil, err
	}

	var elements []types.Node
	for !p.check(types.ARRAY_CLOSE) {
		elem, err := p.parseNode()
		if err != nil {
			return nil, err
		}
		elements = append(elements, elem)

		if p.match(types.COMMA) {
			if p.check(types.ARRAY_CLOSE) {
				return nil, ErrTrailingCommaArray
			}
			continue
		}
		if p.check(types.ARRAY_CLOSE) {
			break
		}
		if p.isAtEnd() {
			return nil, ErrUnexpectedEOF
		}
		return nil, fmt.Errorf("%w: %s at line %d", ErrUnexpectedToken, p.peek().Type, p.peek().Line)
	}

	if _, err := p.consume(types.ARRAY_CLOSE, "expected ']' after array body"); err != nil {
		return nil, err
	}

	return &types.ArrayNode{Elements: elements}, nil
}

func (p *Parser) match(tokenType types.TokenType) bool {
	if p.check(tokenType) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) check(tokenType types.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) consume(tokenType types.TokenType, message string) (types.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	return types.Token{}, fmt.Errorf("%s at line %d", message, p.peek().Line)
}

func (p *Parser) advance() types.Token {
	if !p.isAtEnd() {
		p.pos++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.pos >= len(p.tokens) || p.tokens[p.pos].Type == types.EOF
}

func (p *Parser) peek() types.Token {
	return p.tokens[p.pos]
}

func (p *Parser) previous() types.Token {
	return p.tokens[p.pos-1]
}
