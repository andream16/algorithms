package wigglesort

import "testing"

func TestWiggleSort(t *testing.T) {
	t.Run("it should sort an even number array", func(t *testing.T) {
		arr := []int{3,5,2,1,6,4}
		wiggleSort(arr)
		for i, n := range []int{1, 3, 2, 5, 4, 6} {
			if n != arr[i] {
				t.Fatalf("expected %d at position %d, got %d", n, i, arr[i])
			}
		}
	})
	t.Run("it should sort an odd number array", func(t *testing.T) {
		arr := []int{1, 2, 3, 5, 6}
		wiggleSort(arr)
		for i, n := range []int{1, 3, 2, 6, 5} {
			if n != arr[i] {
				t.Fatalf("expected %d at position %d, got %d", n, i, arr[i])
			}
		}
	})
	t.Run("it should sort an odd number array", func(t *testing.T) {
		arr := []int{1, 2, 3}
		wiggleSort(arr)
		for i, n := range []int{1, 3, 2} {
			if n != arr[i] {
				t.Fatalf("expected %d at position %d, got %d", n, i, arr[i])
			}
		}
	})
}
