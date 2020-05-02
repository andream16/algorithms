package main

import (
	"fmt"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		maxRes, currMax, currMin = nums[0], nums[0], nums[0]
		prevMax, prevMin         = nums[0], nums[0]
	)

	for i := 1; i < len(nums); i++ {
		currMax = max(max(prevMax*nums[i], prevMin*nums[i]), nums[i])
		currMin = min(min(prevMax*nums[i], prevMin*nums[i]), nums[i])
		maxRes = max(maxRes, currMax)
		prevMax = currMax
		prevMin = currMin
	}

	return maxRes
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	//fmt.Println(maxProduct([]int{2, 3, -2, 4}))                     // 6
	//fmt.Println(maxProduct([]int{-2, 0, -1}))                       // 0
	fmt.Println(maxProduct([]int{2, 6, -3, 1, 0, -3, -7, 8, 0, 9})) // 168
	//fmt.Println(maxProduct([]int{-1, -1}))                          // 1
	//fmt.Println(maxProduct([]int{-4, -3, -2}))                      // 12
	//fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))                // 24
}
