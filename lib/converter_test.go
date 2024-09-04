package lib_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/lib"
	"github.com/sarkarshuvojit/lomboktojson/types"
)

func deepEquality(a, b string) bool {
	var ajson map[string]interface{}
	_ = json.Unmarshal([]byte(a), &ajson)
	var bjson map[string]interface{}
	_ = json.Unmarshal([]byte(b), &bjson)

	return reflect.DeepEqual(ajson, bjson)
}

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

func Test_ConvertBasicCustomer(t *testing.T) {
	converter := lib.NewConverter()
	inputTokens := []types.Token{
		{
			Type:    types.CLASS_NAME,
			Lexeme:  "Customer",
			Literal: map[string]string{},
			Line:    0,
		},
		{
			Type:    types.PAREN_OPEN,
			Lexeme:  "(",
			Literal: map[string]string{},
			Line:    0,
		},
		{
			Type:    types.KEY,
			Lexeme:  "name",
			Literal: map[string]string{},
			Line:    0,
		},
		{
			Type:    types.EQUALS,
			Lexeme:  "=",
			Literal: map[string]string{},
			Line:    0,
		},
		{
			Type:    types.VALUE,
			Lexeme:  "Michael Scott",
			Literal: map[string]string{},
			Line:    0,
		},
		{
			Type:    types.PAREN_CLOSE,
			Lexeme:  ")",
			Literal: map[string]string{},
			Line:    0,
		},
		{
			Type:    types.EOF,
			Lexeme:  "",
			Literal: map[string]string{},
			Line:    0,
		},
	}
	t.Run("Should scan simple empty doc", func(t *testing.T) {
		expected := `{"name": "Michael Scott"}`
		got := converter.Convert(inputTokens)

		if !deepEquality(expected, got) {
			t.Errorf("Wanted %s, got %s", expected, got)
		}
	})
}
