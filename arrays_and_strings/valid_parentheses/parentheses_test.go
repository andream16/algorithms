package parentheses

import "testing"

func TestIsValid(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected bool
	}{
		{input: "()", expected: true},
		{input: "()[]{}", expected: true},
		{input: "(]", expected: false},
		{input: "([)]", expected: false},
		{input: "{[]}", expected: true},
	} {
		if valid := isValid(tc.input); tc.expected != valid {
			t.Fatalf("expected %q to be %v, got %v instead", tc.input, tc.expected, valid)
		}
	}
}
