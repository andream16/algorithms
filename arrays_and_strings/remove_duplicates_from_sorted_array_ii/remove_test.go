package remove_duplicates_from_sorted_array_ii

import (
	"fmt"
	"testing"
)

func TestDuplicated(t *testing.T) {
	for i, tt := range []struct {
		givenNums      []int
		expectedNonDup int
		expectedNums   []int
	}{
		{
			givenNums:      []int{1, 1, 1, 2, 2, 3},
			expectedNonDup: 5,
			expectedNums:   []int{1, 1, 2, 2, 3, -1},
		},
		{
			givenNums:      []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expectedNonDup: 7,
			expectedNums:   []int{0, 0, 1, 1, 2, 3, 3, -1, -1},
		},
	} {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			if res := removeDuplicates(tt.givenNums); tt.expectedNonDup != res {
				t.Fatalf("expected %d non duplicates to be removed, got %d", tt.expectedNonDup, res)
			}

			for j := 0; j < len(tt.givenNums); j++ {
				if tt.expectedNums[j] != tt.givenNums[j] {
					t.Fatalf("expected %d at position %d, got %d instead", tt.expectedNums[j], j, tt.givenNums[j])
				}
			}
		})
	}
}
