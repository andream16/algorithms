package meeting

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	for _, tc := range []struct {
		input    [][]int
		expected int
	}{
		{input: [][]int{{7, 10}, {2, 4}}, expected: 1},
		{input: [][]int{{0, 30}, {5, 10}, {15, 20}}, expected: 2},
		{input: [][]int{{1, 10}, {2, 7}, {3, 19}, {8, 12}, {10, 20}, {11, 30}}, expected: 4},
	} {
		if res := minMeetingRooms(tc.input); tc.expected != res {
			t.Fatalf("expected %d, got %d rooms", tc.expected, res)
		}
	}
}
