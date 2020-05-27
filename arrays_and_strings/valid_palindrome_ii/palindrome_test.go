package validpalindromeii

import "testing"

func TestValidPalindrome(t *testing.T) {
	for _, testCase := range []struct{
		input string
		expected bool
	}{
		{input: "aba", expected: true},
		{input: "abca", expected: true},
	}{
		if res := validPalindrome(testCase.input); testCase.expected != res {
			t.Fatalf("expected result for %s to be %v, got %v", testCase.input, testCase.expected, res)
		}
	}
}
