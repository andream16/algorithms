package remove_element

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for i, tt := range []struct {
		givenNums    []int
		target       int
		expectedNums []int
	}{
		{
			givenNums:    []int{3, 2, 2, 3},
			target:       3,
			expectedNums: []int{2, 2, -1, -1},
		},
		{
			givenNums:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			target:       2,
			expectedNums: []int{0, 1, 3, 0, 4, -1, -1, -1},
		},
	} {
		t.Run(fmt.Sprintf("testcase %d", i), func(t *testing.T) {
			removeElement(tt.givenNums, tt.target)
			for j := 0; j < len(tt.givenNums); j++ {
				if tt.expectedNums[j] != tt.givenNums[j] {
					t.Fatalf("expected %d at position %d, got %d instead", tt.expectedNums[j], j, tt.givenNums[j])
				}
			}
		})
	}
}
