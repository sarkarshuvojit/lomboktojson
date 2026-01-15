package scanner

import "errors"

// Sentinel errors for scanner validation.
var (
	ErrKeyExpected   = errors.New("Key is expected")
	ErrValueExpected = errors.New("Value is expected")
)
