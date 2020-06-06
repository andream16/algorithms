package duplicate

// This problem can be solved using a tweak on Floyd's Cycle Detection algorithm.
// When the tortoise meets the hare, we reset the tortoise and slow down the hare to be as fast as the tortoise.
// When they meet again it means that we found the begin of the cycle, and our values.
//
// This is a bit tricky to explain so refer to the explanation on leetcode.
//
// T: O(n)
// S: O(1)
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	slow, fast = nums[slow], nums[nums[fast]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

// Breaks one of the constraints of using constant space.
//
// T: O(n)
// S: O(n)
func findDuplicateNaive(nums []int) int {
	visited := map[int]struct{}{}
	for _, n := range nums {
		if _, ok := visited[n]; ok {
			return n
		}
		visited[n] = struct{}{}
	}
	return -1
}
