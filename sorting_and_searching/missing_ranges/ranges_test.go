package ranges

import "testing"

func TestFindMissingRanges(t *testing.T) {
	rng := findMissingRanges([]int{0, 1, 3, 50, 75}, 0, 99)
	for i, r := range []string{"2","4->49", "51->74", "76->99"} {
		if r != rng[i] {
			t.Fatalf("expected %s, got %s", r, rng[i])
		}
	}
}
