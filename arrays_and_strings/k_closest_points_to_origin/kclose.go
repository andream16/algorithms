package kclose

import (
	"math"
	"sort"
)

// We order by distance and pick the first K element from the sorted slice.
//
// TODO CHECK HEAP AND QUICK SELECT APPROACHES
//
// T: O(n log n)
// S: O(n)
func kClosest(points [][]int, K int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return distance(points[i][0], points[i][1]) < distance(points[j][0], points[j][1])
	})
	return points[:K]
}

func distance(x, y int) float64 {
	xDiff, yDiff := float64(x-0), float64(y-0)
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}
