package maxprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	for _, tc := range []struct {
		input    []int
		expected int
	}{
		{input: []int{7, 1, 5, 3, 6, 4}, expected: 5},
		{input: []int{7, 6, 5, 4, 3, 2, 1}, expected: 0},
	} {
		if res := maxProfit(tc.input); res != tc.expected {
			t.Fatalf("expected %d, got %d instead", tc.expected, res)
		}
	}
}
