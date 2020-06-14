package subarraysum

import "testing"

func TestSubarraySum(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		k        int
		expected int
	}{
		{input: []int{1, 1, 1}, k: 2, expected: 2},
		{input: []int{3, 4, 7, 2, -3, 1, 4, 2}, k: 7, expected: 4},
	} {
		if res := subarraySum(tc.input, tc.k); tc.expected != res {
			t.Fatalf("expected %d, got %d instead", tc.expected, res)
		}
	}
}
