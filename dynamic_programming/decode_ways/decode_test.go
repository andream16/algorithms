package decodeways

import "testing"

func TestNumDecodings(t *testing.T) {
	for _, tc := range []struct{
		input string
		expected int
	}{
		{input: "226", expected: 3},
		{input: "100", expected: 0},
		{input: "101", expected: 1},
		{input: "2264", expected: 3},
	}{
		if ways := numDecodings(tc.input); tc.expected != ways {
			t.Fatalf("expected %d ways, got %d", tc.expected, ways)
		}
	}
}
