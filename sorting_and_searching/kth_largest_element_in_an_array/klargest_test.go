package kthlargest

import "testing"

func TestFindKthLargest(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		k        int
		expected int
	}{
		{input: []int{3, 2, 1, 5, 6, 4}, k: 2, expected: 5},
		{input: []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, k: 4, expected: 4},
	} {
		if n := findKthLargest(tc.input, tc.k); tc.expected != n {
			t.Fatalf("expected %d, got %d", tc.expected, n)
		}
	}
}
