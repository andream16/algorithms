package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	t.Run("testCase1", func(t *testing.T) {
		res := twoSum([]int{2, 7, 11, 15}, 9)
		if 2 != len(res) {
			t.Fatalf("expected 2, got %d indices", len(res))
		}
		if 0 != res[0] {
			t.Fatalf("expected element 0 to be 0, got %d", res[0])
		}
		if 1 != res[1] {
			t.Fatalf("expected element 1 to be 1, got %d", res[1])
		}
	})
	t.Run("testCase2", func(t *testing.T) {
		res := twoSum([]int{2, 7, 19, 20, 11, 15}, 17)
		if 2 != len(res) {
			t.Fatalf("expected 2, got %d indices", len(res))
		}
		if 0 != res[0] {
			t.Fatalf("expected element 0 to be 0, got %d", res[0])
		}
		if 5 != res[1] {
			t.Fatalf("expected element 1 to be 5, got %d", res[1])
		}
	})
	t.Run("testCase3", func(t *testing.T) {
		res := twoSum([]int{-3, 4, 3, 90}, 0)
		if 2 != len(res) {
			t.Fatalf("expected 2, got %d indices", len(res))
		}
		if 0 != res[0] {
			t.Fatalf("expected element 0 to be 0, got %d", res[0])
		}
		if 2 != res[1] {
			t.Fatalf("expected element 1 to be 2, got %d", res[1])
		}
	})
}
