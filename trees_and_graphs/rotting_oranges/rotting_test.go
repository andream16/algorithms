package rotting

import "testing"

func TestOrangesRotting(t *testing.T) {
	t.Run("should return 4 minutes", func(t *testing.T) {
		if mins := orangesRotting([][]int{
			{2, 1, 1},
			{1, 1, 0},
			{0, 1, 1},
		}); mins != 4 {
			t.Fatalf("expected 4 minutes, got %d minutes", mins)
		}
	})
	t.Run("should return -1 because it's not possible to rot one orange", func(t *testing.T) {
		if mins := orangesRotting([][]int{
			{2, 1, 1},
			{0, 1, 1},
			{1, 0, 1},
		}); mins != -1 {
			t.Fatalf("expected -1 because it was not possible to rot one orange, got %d instead", mins)
		}
	})
	t.Run("should return 0 minutes", func(t *testing.T) {
		if mins := orangesRotting([][]int{
			{0, 2},
		}); mins != 0 {
			t.Fatalf("expected 0 minutes, got %d minutes", mins)
		}
	})
	t.Run("should return -1 because it's not possible to rot all the oranges", func(t *testing.T) {
		if mins := orangesRotting([][]int{
			{1},
			{1},
			{1},
			{1},
		}); mins != -1 {
			t.Fatalf("expected -1 because it was not possible to rot all oranges, got %d instead", mins)
		}
	})
	t.Run("should return 0 because there is only one empty cell", func(t *testing.T) {
		if mins := orangesRotting([][]int{
			{0},
		}); mins != 0 {
			t.Fatalf("expected 0 because there is only one empty cell, got %d instead", mins)
		}
	})
}
