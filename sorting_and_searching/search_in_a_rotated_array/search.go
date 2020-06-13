package search

// This problem can be solved by finding the pivot of the array.
// The pivot is the element where the array starts to be rotated.
// Once that's found, we need to use a revised version of binary search to find the target.
//
// A two pass solution is easier to understand. In the first pass, we find the pivot. Then, we check on which subarray
// we should look for the target and, based on that, we binary search for it into it.
//
// A one pass solution, like below, eagerly looks for the subarray.
//
// T: O(log n)
// S: O(1)
func search(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}
	middle := right + left/2
	if target == nums[middle] {
		return middle
	} else if nums[middle] >= nums[left] {
		if target >= nums[left] && target < nums[middle] {
			return binarySearch(nums, left, middle-1, target)
		}
		return binarySearch(nums, middle+1, right, target)
	}
	if target <= nums[right] && target > nums[middle] {
		return binarySearch(nums, middle+1, right, target)
	}
	return binarySearch(nums, left, middle-1, target)
}
