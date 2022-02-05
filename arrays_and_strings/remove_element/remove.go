package remove_element

// T: O(n) assuming that the append operation is constant.
// S: O(n) for the array, the rest is constant.
func removeElement(nums []int, val int) int {
	numsToBeRemoved := 0
	for i := 0; i < len(nums)-numsToBeRemoved; i++ {
		if nums[i] == val {
			numsToBeRemoved++
			nums = append(append(nums[:i], nums[i+1:]...), -1)
			i--
		}
	}
	return len(nums) - numsToBeRemoved
}
