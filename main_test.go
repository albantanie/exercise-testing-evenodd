package main

import "testing"

func TestEvenOdd(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{input: 0, expected: "even"},
		{input: 1, expected: "odd"},
		{input: 2, expected: "even"},
		{input: 3, expected: "odd"},
		{input: 4, expected: "even"},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := evenOdd(test.input)
			if result != test.expected {
				t.Errorf("evenOdd(%d) = %s; want %s", test.input, result, test.expected)
			}
		})
	}
}
