package permutations

import "testing"

func TestPermute(t *testing.T) {
	t.Run("it should correctly permute '[1, 2, 3]'", func(t *testing.T) {
		res := permute([]int{1, 2, 3})
		if 6 != len(res) {
			t.Fatalf("expected 6 permutations, got %d instead", len(res))
		}
	})
}
