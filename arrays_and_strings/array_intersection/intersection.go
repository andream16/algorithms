package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {

	res, tmp := []int{}, []int{}
	set, added := map[int]struct{}{}, map[int]struct{}{}

	if len(nums1) < len(nums2) {
		tmp = nums1
		nums1 = nums2
		nums2 = tmp
	}

	for _, n := range nums1 {
		set[n] = struct{}{}
	}

	for _, n := range nums2 {
		_, ok := set[n]
		if ok {
			_, ok := added[n]
			if ok {
				continue
			}
			res = append(res, n)
			added[n] = struct{}{}
		}
	}

	return res
}

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}
