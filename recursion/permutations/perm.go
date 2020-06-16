package permutations

// To solve this problem we can use recursion.
// If we think about it, a combination is valid when it has the same length of the input.
// So, all we need to do is, have a concept of current combination (left) and left items (right)
// that we need to analyse every time.
// We start with current: [] and left: [1, 2, 3]
// Then we would have current: [1] and left: [2, 3], current: [2] and left: [1, 3], current: [3] and left: [1, 2]
// and we would keep iterating until the base condition is met.
//
// T: O(n!)
// S: O(n!)
func permute(nums []int) [][]int {
	var res [][]int
	permuteRec([]int{}, nums, &res)
	return res
}

func permuteRec(currComb, left []int, res *[][]int) {
	if 0 == len(left) {
		*res = append(*res, currComb)
		return
	}
	for idx, l := range left {
		permuteRec(
			append(currComb, l),
			append(append([]int{}, left[:idx]...), left[idx+1:]...),
			res,
		)
	}
}
