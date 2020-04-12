package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var (
		perms     = [][]int{}
		candidate = []int{}
	)
	if len(nums) == 1 {
		perms = append(perms, nums)
		return perms
	}
	return helper(candidate, nums, perms)
}

func helper(candidate []int, nums []int, res [][]int) [][]int {
	if len(nums) == 0 {
		res = append(append([][]int{}, res...), candidate)
	}
	for i := 0; i < len(nums); i++ {
		candidate = append(candidate, nums[i])
		res = helper(candidate, append(nums[:i], nums[i+1:]...), res)
		candidate = candidate[:len(candidate)-1]
	}
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
