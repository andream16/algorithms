package main

import "fmt"

// T: O(N^2)
// S: constant
func subarraySum(nums []int, k int) int {
	var nSubs, curr int
	if len(nums) == 0 {
		return nSubs
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] == k {
			nSubs++
		}
		curr = nums[i]
		for j := i + 1; j < len(nums); j++ {
			curr += nums[j]
			if curr == k {
				nSubs++
			} else if curr < k {
				continue
			}
		}
	}

	return nSubs
}

func main() {
	fmt.Println(subarraySum([]int{1, 1, 1}, 2))
	fmt.Println(subarraySum([]int{28, 54, 7, -70, 22, 65, -6}, 100))
}
