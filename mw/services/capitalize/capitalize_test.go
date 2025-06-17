package capitalize

import (
	"fmt"
	"testing"
)

func TestCapitalizeCore(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HELLO WORLD"},
		{"go programming", "GO PROGRAMMING"},
		{"capitalize this", "CAPITALIZE THIS"},
		{"", ""},
	}

	for _, test := range tests {
		result := capitalizeCore(test.input)
		fmt.Println(result)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
