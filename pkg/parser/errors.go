package parser

import "errors"

// Sentinel errors for parser validation.
var (
	ErrUnexpectedEOF       = errors.New("parser: unexpected end of input")
	ErrUnexpectedToken     = errors.New("parser: unexpected token")
	ErrTrailingTokens      = errors.New("parser: trailing tokens after parse")
	ErrTrailingCommaObject = errors.New("parser: trailing comma before ')'")
	ErrTrailingCommaArray  = errors.New("parser: trailing comma before ']'")
	ErrMismatchedCloser    = errors.New("parser: mismatched closing token")
)
