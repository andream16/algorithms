package product

import "testing"

func TestProductExceptSelf(t *testing.T) {
	t.Run("should return the expected output", func(t *testing.T) {
		expected := []int{24, 12, 8, 6}
		res := productExceptSelf([]int{1, 2, 3, 4})
		if len(expected) != len(res) {
			t.Fatalf("expected result length to be %d, got %d instead", len(expected), len(res))
		}
		for i := 0; i < len(expected); i++ {
			if expected[i] != res[i] {
				t.Fatalf("expected %d at position %d, got %d instead", expected[i], i, res[i])
			}
		}
	})
}
