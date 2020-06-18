package search

import "testing"

func TestSearchRange(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		target   int
		expected []int
	}{
		{input: []int{5, 7, 7, 8, 8, 10}, target: 8, expected: []int{3, 4}},
		{input: []int{5, 7, 7, 8, 8, 8, 10}, target: 8, expected: []int{3, 5}},
		{input: []int{5, 7, 7, 8, 8, 10}, target: 6, expected: []int{-1, -1}},
		{input: []int{7, 7, 7, 7, 7, 7}, target: 7, expected: []int{0, 5}},
		{input: []int{7, 7}, target: 7, expected: []int{0, 1}},
		{input: []int{7, 7, 8, 9, 10}, target: 8, expected: []int{2, 2}},
		{input: []int{7}, target: 7, expected: []int{0, 0}},
	} {
		res := searchRange(tc.input, tc.target)
		if 2 != len(res) {
			t.Fatalf("expected 2 results, got %d instead", len(res))
		}
		if tc.expected[0] != res[0] || tc.expected[1] != res[1] {
			t.Fatalf("expected range %v, got %v instead", tc.expected, res)
		}
	}
}
