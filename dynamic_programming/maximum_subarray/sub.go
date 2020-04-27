package main

import (
	"fmt"
	"math"
)

// T: O(n)
// S: O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxSum = max(nums[i], maxSum)
	}
	return maxSum
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// T: O(n^2)
// S: 1
func naiveMaxSubArray(nums []int) int {
	maxSum := math.MinInt64

	for i := 0; i < len(nums); i++ {
		currSum := 0
		for j := i; j < len(nums); j++ {
			currSum += nums[j]
			if maxSum < currSum {
				maxSum = currSum
			}
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{1, 2}))                          // 3
	fmt.Println(maxSubArray([]int{-1, 0, -2}))                     // 0
	fmt.Println(maxSubArray([]int{-2, -1}))                        // -1
	fmt.Println(maxSubArray([]int{-1, -2}))                        // -1
	fmt.Println(maxSubArray([]int{-2, -3, -1}))                    // -1
}
