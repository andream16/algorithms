package maximum_product_subarray

// We default all the variables to the first element of the array if it has at least one element.
// From there on, we find the maximum between the current element and the old maximum.
// We do also track the minimum as a negative previous result might translate into a higher maximum if multiplied by
// another negative number.
//
// T: O(n)
// S: O(1)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		maxRes, currMax, currMin = nums[0], nums[0], nums[0]
		prevMax, prevMin         = nums[0], nums[0]
	)

	for i := 1; i < len(nums); i++ {
		currMax = max(max(prevMax*nums[i], prevMin*nums[i]), nums[i])
		currMin = min(min(prevMax*nums[i], prevMin*nums[i]), nums[i])
		maxRes = max(maxRes, currMax)
		prevMax = currMax
		prevMin = currMin
	}

	return maxRes
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
