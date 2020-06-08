package kclose

import (
	"math"
	"sort"
)

type distPoint struct {
	coords   []int
	distance float64
}

// We order by distance and pick the first K element from the sorted slice.
//
// TODO CHECK HEAP AND QUICK SELECT APPROACHES
//
// T: O(n log n)
// S: O(n)
func kClosest(points [][]int, K int) [][]int {
	closest := make([]distPoint, len(points))
	for i, p := range points {
		closest[i] = distPoint{coords: p, distance: distFromOrigin(p[0], p[1])}
	}

	sort.Slice(closest, func(i, j int) bool {
		return closest[i].distance < closest[j].distance
	})

	res := make([][]int, K)
	for i := 0; i < K; i++ {
		res[i] = closest[i].coords
	}

	return res
}

func distFromOrigin(x, y int) float64 {
	xDiff, yDiff := float64(x-0), float64(y-0)
	return math.Sqrt(xDiff*xDiff + yDiff*yDiff)
}
