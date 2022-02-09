package remove_duplicates_from_sorted_array_ii

// If we have less than 3 elements we don't need to do anything.
// Otherwise, we can start with index = 2 and check the previous 2.
// When we find one, we put it at the end of the slice and change its value to -1.
// We must also decrease the index when we find a duplicate otherwise we'd skip the next new elements
// at that position.
// We can stop when we are at index = len(nums)-len(numsToRemove)-1 to avoid infinite loops.
// T: O(n)
// S: O(n)
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	numsToBeRemoved := 0

	for i := 2; i < len(nums)-numsToBeRemoved; i++ {
		if nums[i-2] == nums[i-1] && nums[i-1] == nums[i] {
			numsToBeRemoved++
			i--
			nums = append(append(nums[:i], nums[i+1:]...), -1)
		}
	}

	return len(nums) - numsToBeRemoved
}