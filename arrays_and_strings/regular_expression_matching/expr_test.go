package expr

import "testing"

func TestIsMatch(t *testing.T) {
	for _, testCase := range []struct {
		str     string
		pattern string
		match   bool
	}{
		{str: "aa", pattern: "a", match: false},
		{str: "aa", pattern: "a*", match: true},
		{str: "ab", pattern: ".*", match: true},
		{str: "aab", pattern: "c*a*b", match: true},
		{str: "mississippi", pattern: "mis*is*p*.", match: false},
	} {
		if m := isMatch(testCase.str, testCase.pattern); testCase.match != m {
			t.Fatalf("expected str %s to match pattern %s -> %v", testCase.str, testCase.pattern, testCase.match)
		}
	}
}
