package combination

import "testing"

func TestLetterCombinations(t *testing.T) {
	t.Run("should return the correct combinations for '23'", func(t *testing.T) {
		res := letterCombinations("23")
		if 9 != len(res) {
			t.Fatalf("expected 9 combinations, got %d instead", len(res))
		}
		expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
		for i := 0; i < len(expected); i++ {
			if expected[i] != res[i] {
				t.Fatalf("expected combination %q at position %d, got %q instead", expected[i], i, res[i])
			}
		}
	})
}
