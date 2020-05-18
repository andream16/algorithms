package main

import "testing"

func TestFindAnagrams(t *testing.T) {
	t.Run("testCase1", func(t *testing.T) {
		arr := findAnagrams("cbaebabacd", "abc")
		expectedArr := []int{0, 6}
		if 2 != len(arr) {
			t.Fatalf("expected 2 results, got %d", len(arr))
		}
		for i := 0; i < len(arr); i++ {
			if expectedArr[i] != arr[i] {
				t.Fatalf("expected element %d to be %d, got %d", i, expectedArr[i], arr[i])
			}
		}
	})
	t.Run("testCase2", func(t *testing.T) {
		arr := findAnagrams("abab", "ab")
		expectedArr := []int{0, 1, 2}
		if 3 != len(arr) {
			t.Fatalf("expected 3 results, got %d", len(arr))
		}
		for i := 0; i < len(arr); i++ {
			if expectedArr[i] != arr[i] {
				t.Fatalf("expected element %d to be %d, got %d", i, expectedArr[i], arr[i])
			}
		}
	})
}
