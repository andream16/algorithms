package sub

// To solve this problem, we rely on updating the current value only if its value would increase by adding its previous
// one. On every iteration, we update the maximum subarray sum by checking if the current one is bigger than the
// temporary maximum. This problem is using Kadane's Algorithm.
//
// T: O(n)
// S: O(1)
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		maxSum = max(nums[i], maxSum)
	}
	return maxSum
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
