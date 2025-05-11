package pkg_test

import (
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/pkg"
	"github.com/stretchr/testify/assert"
)

func ptr(s string) *string {
	return &s
}

func TestLombokToJson_NestedValidInputs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Flat structure",
			input:    "Customer(name=Raju,email=raju@gmail.com,age=15)",
			expected: `{"name":"Raju","email":"raju@gmail.com","age":15}`,
		},
		{
			name:     "Nested object",
			input:    "Order(id=123,customer=Customer(name=Raju,email=raju@gmail.com,age=15),amount=500.0)",
			expected: `{"id":123,"customer":{"name":"Raju","email":"raju@gmail.com","age":15},"amount":500.0}`,
		},
		{
			name:     "Nested with multiple objects",
			input:    "Response(status=200,body=Order(id=123,customer=Customer(name=Raju,email=raju@gmail.com),amount=500.0),success=true)",
			expected: `{"status":200,"body":{"id":123,"customer":{"name":"Raju","email":"raju@gmail.com"},"amount":500.0},"success":true}`,
		},
		{
			name:     "Deeply nested",
			input:    "Wrapper(data=Response(status=200,body=Order(id=123,customer=Customer(name=Raju),amount=500.0)))",
			expected: `{"data":{"status":200,"body":{"id":123,"customer":{"name":"Raju"},"amount":500.0}}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := pkg.LombokToJson(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result == nil {
				t.Fatalf("Expected non-nil result for %s", tt.name)
			}
			assert.JSONEq(t, tt.expected, *result)
			/*if *result != tt.expected {
				t.Errorf("For %s:\nExpected: %s\nGot: %s", tt.name, tt.expected, *result)
			}*/
		})
	}
}
