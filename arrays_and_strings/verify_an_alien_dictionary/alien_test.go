package alien

import "testing"

func TestIsAlienSorted(t *testing.T) {
	t.Run("should be sorted", func(t *testing.T) {
		if !isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz") {
			t.Fatal("should have been sorted but it's not")
		}
	})
	t.Run("should be sorted", func(t *testing.T) {
		if !isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz") {
			t.Fatal("should have been sorted but it's not")
		}
	})
	t.Run("should not be sorted", func(t *testing.T) {
		if isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz") {
			t.Fatal("should have not been sorted but it's")
		}
	})
	t.Run("should not be sorted", func(t *testing.T) {
		if isAlienSorted([]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz") {
			t.Fatal("should have not been sorted but it's")
		}
	})
}
