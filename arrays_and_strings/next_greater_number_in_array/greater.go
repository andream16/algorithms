package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, len(nums1))
	if len(nums1) == 0 {
		return res
	}

	m := map[int][]int{}

	for idx, n := range nums2 {
		idx++
		m[n] = nums2[idx:]
	}

	for idx, n := range nums1 {
		next, ok := m[n]
		res[idx] = -1
		if !ok {
			continue
		}
		for _, nxt := range next {
			if nxt > n {
				res[idx] = nxt
				break
			}
		}
	}

	return res

}

// Input: nums1 = [4,1,2], nums2 = [1,3,4,2].
// Output: [-1,3,-1]

func main() {
	//nums1 := []int{4, 1, 2}
	//nums2 := []int{1, 3, 4, 2}
	nums1 := []int{}
	nums2 := []int{1, 2, 3, 4, 5}
	fmt.Println(nextGreaterElement(nums1, nums2))
}
