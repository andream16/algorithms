package permutations

// T: O(N!)
// S: O(N!)
func permute(nums []int) [][]int {
	var perms [][]int
	if len(nums) == 1 {
		perms = append(perms, nums)
		return perms
	}
	backtrack(0, &nums, &perms)
	return perms
}

// by moving the index & swapping current with next and back, we get all permutations starting with a given index first.
func backtrack(first int, nums *[]int, res *[][]int) {
	if first == len(*nums) {
		*res = append(*res, append([]int{}, *(nums)...))
	}
	for i := first; i < len(*nums); i++ {
		(*nums)[first], (*nums)[i] = (*nums)[i], (*nums)[first]
		backtrack(first+1, nums, res)
		(*nums)[first], (*nums)[i] = (*nums)[i], (*nums)[first]
	}
}
