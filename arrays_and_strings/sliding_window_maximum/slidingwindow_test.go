package slidingwindow

import "testing"

func TestMaxSlidingWindow(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		k        int
		expected []int
	}{
		{input: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3, expected: []int{3, 3, 5, 5, 6, 7}},
		{input: []int{1}, k: 1, expected: []int{1}},
		{input: []int{1, 2, 3}, k: 1, expected: []int{1, 2, 3}},
	} {
		res := maxSlidingWindow(tc.input, tc.k)
		if len(tc.expected) != len(res) {
			t.Fatalf("expected %d results, got %d instead", len(tc.expected), len(res))
		}
		for i := 0; i < len(tc.expected); i++ {
			if tc.expected[i] != res[i] {
				t.Fatalf("expected %d at position %d, got %d instead", tc.expected[i], i, res[i])
			}
		}
	}
}
