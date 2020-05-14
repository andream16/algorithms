package kdigits

import "testing"

func TestRemoveKDigits(t *testing.T) {
	for _, testCase := range []struct {
		givenNum string
		givenK   int
		expected string
	}{
		{givenNum: "12345264", givenK: 4, expected: "1224"},
		{givenNum: "1432219", givenK: 3, expected: "1219"},
		{givenNum: "10200", givenK: 1, expected: "200"},
		{givenNum: "10", givenK: 2, expected: "0"},
		{givenNum: "10", givenK: 1, expected: "0"},
		{givenNum: "112", givenK: 1, expected: "11"},
	} {
		if res := removeKdigits(testCase.givenNum, testCase.givenK); res != testCase.expected {
			t.Fatalf("expected %q, got %q", testCase.expected, res)
		}
	}
}
