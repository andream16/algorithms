package main

import (
	"fmt"
	"sort"
)

// T: O(N*log(N)) for sorting. Rest is constant.
// S: N1+N2
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	l := len(nums1)
	n1, n2 := nums1[l/2], 0
	if l > 1 {
		n2 = nums1[(l/2)-1]
	}
	if l%2 != 0 {
		return float64(n1)
	}
	return float64(n1+n2) / 2
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))    // 2
	fmt.Println(findMedianSortedArrays([]int{}, []int{1}))        // 1
}
