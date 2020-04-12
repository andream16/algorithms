package main

import "fmt"

// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2,2]

// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [4,9]

func intersect(nums1 []int, nums2 []int) []int {

	if len(nums1) > len(nums2) {
		tmp := nums1
		nums1 = nums2
		nums2 = tmp
	}

	m := make(map[int]int, len(nums2))
	for _, n := range nums2 {
		c, ok := m[n]
		if !ok {
			m[n] = 1
			continue
		}
		m[n] = c + 1
	}

	var res []int

	for _, n := range nums1 {
		c, ok := m[n]
		if ok && c > 0 {
			m[n] = c - 1
			res = append(res, n)
		}
	}

	return res

}

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))       // [2,2]
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})) // [4,9]
	fmt.Println(intersect([]int{3, 1, 2}, []int{1, 1}))          // [1]
}
