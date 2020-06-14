package sub

import "testing"

func TestMaxSubArray(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		expected int
	}{
		{input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expected: 6},
		{input: []int{1}, expected: 1},
		{input: []int{1, 2}, expected: 3},
		{input: []int{-1, 0, -2}, expected: 0},
		{input: []int{-2, -1}, expected: -1},
		{input: []int{-1, -2}, expected: -1},
		{input: []int{-2, -3, -1}, expected: -1},
	} {
		if res := maxSubArray(tc.input); tc.expected != res {
			t.Fatalf("expected %d, got %d for array %v", tc.expected, res, tc.input)
		}
	}
}
