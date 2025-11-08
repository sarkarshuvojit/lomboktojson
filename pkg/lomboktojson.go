package pkg

import (
	"bytes"
	"errors"

	"github.com/sarkarshuvojit/lomboktojson/pkg/beautifier"
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

// Beautify returns the formatted representation of the provided Lombok toString
// output using the beautifier package. If indent is less than or equal to zero,
// the beautifier will default to single-space indentation.
func Beautify(in string, indent int) (string, error) {
	return beautifier.BeautifySource(in, indent)
}
