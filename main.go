package main

import (
	"bytes"
	"fmt"

	"github.com/sarkarshuvojit/lomboktojson/lib"
)

func main() {
	source := `Customer()`
	var sourceBuf bytes.Buffer
	sourceBuf.WriteString(source)

	scanner := lib.NewScanner(&sourceBuf)
	tokens := scanner.Scan()
	if val, err := (lib.Convert(tokens)); err != nil {
		fmt.Println(string(val))
	}
}
