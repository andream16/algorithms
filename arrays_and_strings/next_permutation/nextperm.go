package main

import (
	"fmt"
	"math"
)

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

func main() {
	a, b, c, d, e := []int{1, 2, 3}, []int{3, 2, 1}, []int{1, 1, 5}, []int{1, 5, 8, 4, 7, 6, 5, 3, 1}, []int{0, 1, 1, 0, 4, 4}
	f := []int{100, 99, 98, 97, 96, 95, 94, 93, 92, 91, 90, 89, 88, 87, 86, 85, 84, 83, 82, 81, 80, 79, 78, 77, 76, 75, 74, 73, 72, 71, 70, 69, 68, 67, 66, 65, 64, 63, 62, 61, 60, 59, 58, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 47, 46, 45, 44, 43, 42, 41, 40, 39, 38, 37, 36, 35, 34, 33, 32, 31, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	nextPermutation(a)
	nextPermutation(b)
	nextPermutation(c)
	nextPermutation(d)
	nextPermutation(e)
	nextPermutation(f)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
