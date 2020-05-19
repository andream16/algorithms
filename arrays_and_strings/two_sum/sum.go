package twosum

// We save for every element nums[i], target - nums[i].
// So, every time we process a new element, we check if it's the one missing to reach the target.
//
// T: O(n)
// S: O(n)
func twoSum(nums []int, target int) []int {

	var (
		res = make([]int, 2)
		m   = make(map[int]int, len(nums))
	)

	for i := 0; i < len(nums); i++ {
		if idx, ok := m[nums[i]]; ok {
			res[0], res[1] = idx, i
			break
		}
		m[target-nums[i]] = i
	}

	return res
}
