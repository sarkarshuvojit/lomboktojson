package pkg_test

import (
	"errors"
	"testing"

	"github.com/sarkarshuvojit/lomboktojson/pkg"
	"github.com/sarkarshuvojit/lomboktojson/pkg/scanner"
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
		{
			name:     "Nested array of objects",
			input:    "Invoice(number=INV1234, date=2023-12-01, items=[Item(id=1),Item(id=2)])",
			expected: `{"number":"INV1234","date":"2023-12-01","items":[{"id":1},{"id":2}]}`,
		},
		{
			name:     "Simple array of numbers",
			input:    "Scores(values=[90,85,95])",
			expected: `{"values":[90,85,95]}`,
		},
		{
			name:     "Simple array of strings",
			input:    "Basket(items=[apple,banana,orange])",
			expected: `{"items":["apple","banana","orange"]}`,
		},
		{
			name:     "Array of mixed literals",
			input:    "Flags(values=[true,false,null,0,1])",
			expected: `{"values":[true,false,null,0,1]}`,
		},
		{
			name:     "Nested arrays",
			input:    "Matrix(values=[[1,2],[3,4]])",
			expected: `{"values":[[1,2],[3,4]]}`,
		},
		{
			name:     "Truncated decimal treated as string",
			input:    "Product(id=102,name=Laptop,price=[999.99...],inStock=true)",
			expected: `{"id":102,"name":"Laptop","price":["999.99..."],"inStock":true}`,
		},
		{
			name:     "Non-numeric float-like string",
			input:    "Product(id=102,name=Laptop,price=99.99ggwwp,inStock=true)",
			expected: `{"id":102,"name":"Laptop","price":"99.99ggwwp","inStock":true}`,
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

func TestLombokToJson_InvalidInputMissingKey(t *testing.T) {
	input := "Product(id=1,=something)"
	result, err := pkg.LombokToJson(input)
	if err == nil {
		t.Fatalf("Expected error but got none with result: %v", result)
	}
	if !errors.Is(err, scanner.ErrKeyExpected) {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestLombokToJson_InvalidInputMissingValue(t *testing.T) {
	input := "Product(id=1,name=,inStock=true)"
	result, err := pkg.LombokToJson(input)
	if err == nil {
		t.Fatalf("Expected error but got none with result: %v", result)
	}
	if !errors.Is(err, scanner.ErrValueExpected) {
		t.Fatalf("Unexpected error: %v", err)
	}
}
