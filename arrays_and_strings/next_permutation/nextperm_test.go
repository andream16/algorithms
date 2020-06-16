package nextperm

import "testing"

func TestNextPermutation(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		expected []int
	}{
		{input: []int{1, 2, 3}, expected: []int{1, 3, 2}},
		{input: []int{3, 2, 1}, expected: []int{1, 2, 3}},
		{input: []int{1, 1, 5}, expected: []int{1, 5, 1}},
	} {
		nextPermutation(tc.input)
		for i := 0; i < len(tc.expected); i++ {
			if tc.expected[i] != tc.input[i] {
				t.Fatalf("expected %d at position %d, got %d instead", tc.expected[i], i, tc.input[i])
			}
		}
	}
}
