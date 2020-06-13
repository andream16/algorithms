package search

import "testing"

func TestSearch(t *testing.T) {
	for _, tc := range []struct {
		nums     []int
		target   int
		expected int
	}{
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expected: 4},
		{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, expected: 4},
		{nums: []int{}, target: 5, expected: -1},
		{nums: []int{1}, target: 1, expected: 0},
		{nums: []int{1, 3}, target: 3, expected: 1},
	} {
		if idx := search(tc.nums, tc.target); tc.expected != idx {
			t.Fatalf("expected %d, got %d", tc.expected, idx)
		}
	}
}
