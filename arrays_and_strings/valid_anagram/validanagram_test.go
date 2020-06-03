package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	for _, tc := range []struct{
		input [2]string
		expected bool
	}{
		{input: [2]string{"anagram", "nagaram"}, expected: true},
		{input: [2]string{"rat", "car"}, expected: false},
	}{
		if res := isAnagram(tc.input[0], tc.input[1]); tc.expected != res {
			t.Fatalf("expected %v, bot got %v for strings %q and %q", tc.expected, res, tc.input[0], tc.input[1])
		}
	}
}
