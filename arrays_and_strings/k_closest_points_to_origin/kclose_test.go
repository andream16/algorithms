package kclose

import "testing"

func TestKClosest(t *testing.T) {
	t.Run("should return what expected", func(t *testing.T) {
		expected := [][]int{{3, 3}, {-2, 4}}
		res := kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)
		if 2 != len(res) {
			t.Fatalf("expected 2 points, got %d instead", len(res))
		}
		for i := 0; i < len(res); i++ {
			if expected[i][0] != res[i][0] || expected[i][1] != res[i][1] {
				t.Fatalf(
					"expected point [%d][%d] at position %d, got [%d][%d] instead",
					expected[i][0],
					expected[i][1],
					i,
					res[i][0],
					res[i][1],
				)
			}
		}
	})
}
