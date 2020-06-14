package longest

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, tc := range []struct {
		input    string
		expected int
	}{
		{input: "abcabcbb", expected: 3},
		{input: "bbbbb", expected: 1},
		{input: "pwwkew", expected: 3},
	} {
		if res := lengthOfLongestSubstring(tc.input); tc.expected != res {
			t.Fatalf("expected %d for string %q, got %d instead", tc.expected, tc.input, res)
		}
	}
}
