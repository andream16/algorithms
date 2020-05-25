package mergesortedarrays

import "testing"

func TestMerge(t *testing.T) {
	t.Run("should merge array", func(t *testing.T) {
		input1, input2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
		expected := []int{1, 2, 2, 3, 5, 6}
		merge(input1, 3, input2, 3)
		for i := 0; i < len(input1); i++ {
			if expected[i] != input1[i] {
				t.Fatalf("expected %d at position %d, got %d", expected[i], i, input1[i])
			}
		}
	})
}
