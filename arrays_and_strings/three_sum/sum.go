package sum

import (
	"fmt"
	"math"
	"sort"
)

// If we sort the elements, using a binary search approach, we can binary search for the values that will satisfy the
// condition: complement == (left+right). We use a set to understand if we visited a triplet or not.
//
// TODO: revisit
//
// T: O(n^2)
// S: O(n^2)
func threeSum(nums []int) [][]int {

	var (
		res  [][]int
		seen = map[string]struct{}{}
		saw  = func(x, y, z int) bool {
			s := fmt.Sprintf("%d%d%d", x, y, z)
			_, ok := seen[s]
			if ok {
				return true
			}
			seen[s] = struct{}{}
			return false
		}
		left, right     = math.MinInt64, 0
		complement, sum int
	)

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		complement = 0 - nums[i]
		left = i + 1
		right = len(nums) - 1
		for left < right {
			sum = nums[left] + nums[right]
			if complement == sum {
				if !saw(nums[i], nums[left], nums[right]) {
					res = append(res, []int{nums[i], nums[left], nums[right]})
				}
			}
			if sum > complement {
				right--
				continue
			}
			left++
		}
	}

	return res

}
