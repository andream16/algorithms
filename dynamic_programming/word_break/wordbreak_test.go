package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	for _, tc := range []struct{
		inputStr string
		inputDict []string
		expected bool
	}{
		{inputStr: "leetcode", inputDict: []string{"leet", "code"}, expected: true},
		{inputStr: "applepenapple", inputDict: []string{"apple", "pen"}, expected: true},
		{inputStr: "catsandog", inputDict: []string{"cats", "dog", "sand", "and", "cat"}, expected: false},
	}{
		if res := wordBreak(tc.inputStr, tc.inputDict); tc.expected != res {
			t.Fatalf("expected %v, got %v for string %s", tc.expected, res, tc.inputStr)
		}
	}
}
