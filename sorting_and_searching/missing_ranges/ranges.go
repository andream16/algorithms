package ranges

import "fmt"

// TODO CHECK THIS
func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		if lower == upper {
			return []string{fmt.Sprintf("%d", lower)}
		}
		return []string{fmt.Sprintf("%d->%d", lower, upper)}
	}

	var rng []string
	for i := 1; i < len(nums); i++ {
		low, high := nums[i-1]+1, nums[i]-1
		if low == high {
			rng = append(rng, fmt.Sprintf("%d", low))
		} else if high > low {
			rng = append(rng, fmt.Sprintf("%d->%d", low, high))
		}
	}
	if lower < nums[0]-1 {
		rng = append([]string{fmt.Sprintf("%d->%d", lower, nums[0]-1)}, rng...)
	}
	if upper > nums[len(nums)-1]+1 {
		rng = append(rng, fmt.Sprintf("%d->%d", nums[len(nums)-1]+1, upper))
	}
	return rng
}
