package gameoflife

import "testing"

func TestGameOfLife(t *testing.T) {
	t.Run("should return the expected values", func(t *testing.T) {
		input := [][]int{
			{0, 1, 0},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 0, 0},
			{1, 0, 1},
			{0, 1, 1},
			{0, 1, 0},
		}
		gameOfLife(input)
		for i := 0; i < len(expected); i++ {
			for j := 0; j < len(expected[i]); j++ {
				if expected[i][j] != input[i][j] {
					t.Fatalf(
						"expected %d at position (%d,%d), got %d instead",
						expected[i][j],
						i,
						j,
						input[i][j],
					)
				}
			}
		}
	})
}
