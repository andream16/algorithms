package main

import (
	"fmt"
	"math"
)

// TODO improve

// T: O(N-1*((N*N-1)*N)
// S: O(N)
func merge(intervals [][]int) [][]int {

	var (
		curr  int
		found bool
	)

	for curr < len(intervals)-1 {
		currMin, currMax := findMinMax(intervals[curr])
		for i := curr + 1; i < len(intervals); i++ {
			nextMin, nextMax := findMinMax(intervals[i])
			if min, max, overlap := overlaps(currMin, nextMin, currMax, nextMax); overlap {
				intervals[curr] = []int{min, max}
				intervals = append(intervals[:i], intervals[i+1:]...)
				found = true
				break
			}
		}
		if found {
			found = false
			curr = 0
		} else {
			curr++
		}

	}

	return intervals

}

func overlaps(min1, min2, max1, max2 int) (int, int, bool) {
	min, max := findMinMax([]int{min1, min2, max1, max2})
	return min, max, min1 <= max2 && min2 <= max1
}

func findMinMax(arr []int) (int, int) {
	min, max := math.MaxInt64, math.MinInt64
	for _, n := range arr {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	fmt.Println(merge([][]int{{1, 4}, {4, 5}}))
	fmt.Println(merge([][]int{{1, 4}, {0, 1}}))
}
