package mergeintervals

import (
	"sort"
)

// We sort the intervals by their beginning. We add the first interval to the result array and we iterate the rest.
// Since both input and result are sorted, we can merge the current interval with the last result if current interval
// begins before the end of the last result. In the other cases we append a new interval to the result.
//
// T: O(n log n)
// S: O(n)
func merge(intervals [][]int) [][]int {

	if 0 == len(intervals) {
		return make([][]int, 0)
	}

	var res [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastResIdx := len(res) - 1
		if intervals[i][0] > res[lastResIdx][1] {
			res = append(res, intervals[i])
		} else if res[lastResIdx][1] < intervals[i][1] {
			res[lastResIdx][1] = intervals[i][1]
		}
	}

	return res
}
