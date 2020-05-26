package palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, testCase := range []struct {
		input    int
		expected bool
	}{
		{input: 121, expected: true},
		{input: 100, expected: false},
		{input: 0, expected: true},
		{input: 22, expected: true},
	} {
		if res := isPalindrome(testCase.input); testCase.expected != res {
			t.Fatalf("expected input %d to be palindromic == %v, got %v", testCase.input, testCase.expected, res)
		}
	}
}
