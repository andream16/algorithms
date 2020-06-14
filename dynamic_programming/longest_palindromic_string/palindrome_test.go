package longestpalindrome

import "testing"

func TestLongestPalindrome(t *testing.T) {
	for _, tc := range []struct {
		input  string
		output string
	}{
		{input: "babad", output: "bab"},
		{input: "cbbd", output: "bb"},
	} {
		if p := longestPalindrome(tc.input); tc.output != p {
			t.Fatalf("expected %q, got %q instead", tc.output, p)
		}
	}
}
