package lib

import "bytes"

func LombokToJson(in string) string {
	var sourceBuf bytes.Buffer
	sourceBuf.WriteString(in)
	scanner := NewScanner(&sourceBuf)
	tokens := scanner.Scan()

	converter := NewConverter()
	return converter.Convert(tokens)
}
