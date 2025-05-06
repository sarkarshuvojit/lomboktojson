package pkg

import (
	"bytes"
	"errors"

	"github.com/sarkarshuvojit/lomboktojson/pkg/generator"
	"github.com/sarkarshuvojit/lomboktojson/pkg/scanner"
)

// LombokToJson converts the output of Lombok's default toString() method
// into a well-structured JSON string.
//
// This utility is helpful when dealing with logs or outputs that include
// Lombok-formatted toString() results, making them easier to parse or process.
//
// Returns a pointer to a JSON string on success, or an error if conversion fails.
func LombokToJson(in string) (*string, error) {
	var sourceBuf bytes.Buffer
	sourceBuf.WriteString(in)

	scanner := scanner.NewScanner(&sourceBuf)
	tokens := scanner.Scan()
	if val, err := generator.Generate(tokens); err == nil {
		valStr := string(val)
		return &valStr, nil
	} else {
		return nil, errors.New("ERROR: Unable to convert at the moment")
	}
}
