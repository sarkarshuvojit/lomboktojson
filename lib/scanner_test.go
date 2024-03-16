package lib_test

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/lib"
)

func Test_ScanSimple(t *testing.T) {
	t.Run("Should scan simple empty doc", func(t *testing.T) {
		source := ``
		tokens := lib.Scan(source)
		expectedTokenLen := 1
		if len(tokens) != expectedTokenLen {
			t.Errorf("Should return one token for EOF, got %d", len(tokens))
		}
	})
}
