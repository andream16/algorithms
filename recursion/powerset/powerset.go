package main

import "fmt"

// T: O(N*2^N) to generate and copy
// S: O(N*2^N) to generate and copy
func subsets(nums []int) [][]int {
	return helper(nums, 0, [][]int{{}})
}

func helper(nums []int, i int, sets [][]int) [][]int {
	if i == len(nums) {
		return sets
	}
	for _, s := range sets {
		sets = append(sets, append([]int{nums[i]}, s...))
	}
	return helper(nums, i+1, sets)
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
