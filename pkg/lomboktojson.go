package pkg

import (
	"bytes"
	"errors"

	"github.com/sarkarshuvojit/lomboktojson/pkg/generator"
	"github.com/sarkarshuvojit/lomboktojson/pkg/scanner"
)

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
