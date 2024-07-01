package lib_test

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/lib"
	"github.com/sarkarshuvojit/lomboktojson/types"
)

func Test_ScanSimpleDocument(t *testing.T) {
	converter := lib.NewConverter()
	inputTokens := []types.Token{
		{
			Type:    types.EOF,
			Lexeme:  "",
			Literal: map[string]string{},
			Line:    0,
		},
	}
	t.Run("Should scan simple empty doc", func(t *testing.T) {
		expected := `{}`
		got := converter.Convert(inputTokens)
		if expected != got {
			t.Errorf("Wanted %s, got %s", expected, got)
		}
	})
}
