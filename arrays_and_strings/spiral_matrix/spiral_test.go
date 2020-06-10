package spiral

import "testing"

func TestSpiralOrder(t *testing.T) {
	t.Run("should return a slice with the correct spiral order of a squared matrix", func(t *testing.T) {
		expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		order := spiralOrder([][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		})
		if len(expected) != len(order) {
			t.Fatalf("expected %d elements, got %d instead", len(expected), len(order))
		}
		for i := 0; i < len(order); i++ {
			if expected[i] != order[i] {
				t.Fatalf("expected %d at position %d, got %d instead", expected[i], i, order[i])
			}
		}
	})
	t.Run("should return a slice with the correct spiral order of a non squared matrix", func(t *testing.T) {
		expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		order := spiralOrder([][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		})
		if len(expected) != len(order) {
			t.Fatalf("expected %d elements, got %d instead", len(expected), len(order))
		}
		for i := 0; i < len(order); i++ {
			if expected[i] != order[i] {
				t.Fatalf("expected %d at position %d, got %d instead", expected[i], i, order[i])
			}
		}
	})
	t.Run("should return an empty slice for an empty matrix", func(t *testing.T) {
		var expected []int
		order := spiralOrder([][]int{})
		if len(expected) != len(order) {
			t.Fatalf("expected %d elements, got %d instead", len(expected), len(order))
		}
	})
}
