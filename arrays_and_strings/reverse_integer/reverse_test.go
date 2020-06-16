package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		input    int
		expected int
	}{
		{input: 321, expected: 123},
		{input: -123, expected: -321},
		{input: 120, expected: 21},
	} {
		if res := reverse(tc.input); tc.expected != res {
			t.Fatalf("expected %d, got %d instead for input %d", tc.expected, res, tc.input)
		}
	}
}
