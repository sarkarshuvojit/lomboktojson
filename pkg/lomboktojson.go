package pkg

import (
	"bytes"
	"errors"

	"github.com/sarkarshuvojit/lomboktojson/lib"
)

func LombokToJson(in string) (*string, error) {
	var sourceBuf bytes.Buffer
	sourceBuf.WriteString(in)

	scanner := lib.NewScanner(&sourceBuf)
	tokens := scanner.Scan()
	if val, err := lib.Generate(tokens); err == nil {
		valStr := string(val)
		return &valStr, nil
	} else {
		return nil, errors.New("ERROR: Unable to convert at the moment")
	}
}
