package maximum_product_subarray

import "testing"

func TestMaxProducts(t *testing.T) {
	for _, tt := range []struct {
		nums        []int
		expectedMax int
	}{
		{
			nums:        []int{2, 6, -3, 1, 0, -3, -7, 8, 0, 9},
			expectedMax: 168,
		},
		{
			nums:        []int{-2},
			expectedMax: -2,
		},
		{
			nums:        []int{},
			expectedMax: 0,
		},
	} {
		if res := maxProduct(tt.nums); res != tt.expectedMax {
			t.Fatalf("expected product for %v to be %d, got %d instead", tt.nums, tt.expectedMax, res)
		}
	}
}
