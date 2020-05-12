package main

import "fmt"

// We binary search for the first match. If no match is found, we return -1 and then [-1, -1].
// If we find a match, we decrease low and increase high indexes until the element at their
// position has their same value as the target. We stop when they are both different or we reach
// the start and end of the array (in case all the elements have the same value).
//
// T: O(log n)
// S: O(1)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	firstIdx := binarySearchFirstIdx(0, len(nums)-1, target, nums)
	if -1 == firstIdx {
		return []int{-1,-1}
	}

	low, high := firstIdx, firstIdx

	for (low > 0 && target == nums[low-1]) || (high < len(nums)-1 && target == nums[high+1]) {
		if low > 0 && target == nums[low-1] {
			low = low-1
		}
		if high < len(nums)-1 && target == nums[high+1] {
			high = high+1
		}
	}

	return []int{low,high}

}

func binarySearchFirstIdx(start, end, target int, arr []int) int {
	if start == end {
		if arr[start] == arr[end] && arr[end] == target {
			return start
		}
		return -1
	}
	mid := (start+end) /2
	if arr[mid] < target {
		return binarySearchFirstIdx(mid+1, end, target, arr)
	}
	if arr[mid] > target {
		return binarySearchFirstIdx(start, mid, target, arr)
	}
	return mid
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8)) // [3,4]
	fmt.Println(searchRange([]int{5,7,7,8,8,8, 10}, 8)) // [3,5]
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 6)) // [-1,-1]
	fmt.Println(searchRange([]int{7,7,7,7,7,7}, 7)) // [0, 5]
	fmt.Println(searchRange([]int{7, 7}, 7)) // [0, 1]
	fmt.Println(searchRange([]int{7, 7, 8, 9, 10}, 8)) // [2,2]
	fmt.Println(searchRange([]int{7}, 7)) // [0, 0]
}
