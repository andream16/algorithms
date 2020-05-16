package average

type MovingAverage struct {
	maxSize int
	nums []int
	currSum int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		maxSize: size,
		nums: []int{},
	}
}

// T: O(1)
// S: O(n)
func (this *MovingAverage) Next(val int) float64 {
	if this.maxSize == len(this.nums) {
		this.currSum -= this.nums[0]
		this.currSum += val
		this.nums = append(this.nums[1:], val)
	} else {
		this.currSum += val
		this.nums = append(this.nums, val)
	}
	return float64(this.currSum) / float64(len(this.nums))
}
