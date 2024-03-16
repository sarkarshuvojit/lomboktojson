package lib_test

import (
	"bytes"
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/lib"
	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Test_ScanSimple(t *testing.T) {
	t.Run("Should scan simple empty doc", func(t *testing.T) {
		source := ``
		var sourceBuf bytes.Buffer
		sourceBuf.WriteString(source)

		tokens := lib.Scan(&sourceBuf)

		expectedTokenLen := 1
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return one token for EOF, got %d", len(tokens))
		}

		expectedTokenType := types.EOF
		if tokens[0].Type != expectedTokenType {
			t.Errorf("Should return EOF, got %d", tokens[0].Type)
		}
	})
}
