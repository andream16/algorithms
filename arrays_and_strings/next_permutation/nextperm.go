package nextperm

import (
	"math"
)

// To solve this problem we have to really understand what we have to do.
// Given '123' we need to find the next lexicographically greater permutation. In this case it means '132'.
// To do so, we look for the first element where nums[i] > nums[i-1], starting from the end of the array.
// If we found it, we start looking for the first element that is just greater than the decreasing one.
// If we found it, we swap these two and we reverse the elements after the first index.
// If we didn't, we just need to reverse in the lowest possible order.
//
// T: O(n)
// S: O(n)
func nextPermutation(nums []int) {

	var (
		min                = math.MaxInt64
		candidate          = 0
		firstDecreasingIdx = 0
		minIdx             = 0
		rev                []int
		found              = false
	)

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			candidate = nums[i-1]
			firstDecreasingIdx = i - 1
			found = true
			break
		}
	}

	if !found {
		reverse(nums)
		return
	}

	for i := firstDecreasingIdx; i < len(nums); i++ {
		diff := nums[i] - candidate
		if diff > 0 && diff <= min {
			min = nums[i]
			minIdx = i
		}
	}
	nums[firstDecreasingIdx] = min
	nums[minIdx] = candidate

	for i := len(nums) - 1; i > firstDecreasingIdx; i-- {
		rev = append(rev, nums[i])
	}

	nums = nums[:firstDecreasingIdx+1]
	nums = append(nums, rev...)
}

func reverse(nums []int) {
	tmp, ctr := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		if ctr >= i {
			break
		}
		tmp = nums[i]
		nums[i] = nums[ctr]
		nums[ctr] = tmp
		ctr++
	}
}
