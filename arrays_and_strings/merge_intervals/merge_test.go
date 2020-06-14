package mergeintervals

import "testing"

func TestMerge(t *testing.T) {
	for _, tc := range []struct {
		input    [][]int
		expected [][]int
	}{
		{input: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, expected: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{input: [][]int{{1, 4}, {4, 5}}, expected: [][]int{{1, 5}}},
	} {
		res := merge(tc.input)
		if len(tc.expected) != len(res) {
			t.Fatalf("expected result to have length %d, got %d instead", len(tc.expected), len(res))
		}
		for i := range tc.expected {
			if tc.expected[i][0] != res[i][0] || tc.expected[i][1] != res[i][1] {
				t.Fatalf("expected interval to be %v, got %v instead", tc.expected[i], res[i])
			}
		}
	}
}
