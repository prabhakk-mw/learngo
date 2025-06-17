package capitalize

import (
	"testing"
)

func TestCapitalizeCore(test *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"go programming", "GO PROGRAMMING"},
		{"capitalize this", "CAPITALIZE THIS"},
		{"", ""},
	}

	for _, testpoint := range tests {
		result := capitalizeCore(testpoint.input)
		test.Logf("capitalizeCore(%q) => %q; expected %q", testpoint.input, result, testpoint.expected)

		if result != testpoint.expected {
			test.Errorf("Capitalize(%q) = %q; want %q", testpoint.input, result, testpoint.expected)
		}
	}
}
