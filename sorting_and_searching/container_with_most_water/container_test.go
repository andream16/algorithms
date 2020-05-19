package container

import "testing"

func TestMaxArea(t *testing.T) {
	t.Run("testCase1", func(t *testing.T) {
		if area := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}); 49 != area {
			t.Fatalf("expected area to be 49, got %d", area)
		}
	})
	t.Run("testCase2", func(t *testing.T) {
		if area := maxArea([]int{1, 1}); 1 != area {
			t.Fatalf("expected area to be 1, got %d", area)
		}
	})
	t.Run("testCase3", func(t *testing.T) {
		if area := maxArea([]int{1, 2, 1}); 2 != area {
			t.Fatalf("expected area to be 2, got %d", area)
		}
	})
}
