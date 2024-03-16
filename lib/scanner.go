package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Scan(source io.Reader) (tokens []types.Token) {
	var curline int8
	curline = 1
	sourceBytes, err := io.ReadAll(source)
	if err != nil {
		fmt.Printf("Could not read %v: %v\n", source, err)
		os.Exit(1)
	}

	for ch := range sourceBytes {
		fmt.Println(ch)
	}
	tokens = append(tokens, types.NewToken(types.EOF, "", nil, curline))
	return
}
