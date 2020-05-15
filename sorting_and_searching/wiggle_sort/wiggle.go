package wigglesort

import "sort"

// We sort the array and then we simply swap elements next to each other to satisfy the condition.
//
// T: O(n log n)
// S: O(1)
func wiggleSort(nums []int)  {
	sort.Ints(nums)

	for i:=1; i<len(nums)-1; i+=2 {
		tmp := nums[i]
		nums[i] = nums[i+1]
		nums[i+1] = tmp
	}

}
