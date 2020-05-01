package main

import (
	"fmt"
	"math"
)

// TODO improve

// T: O(N-1*((N*N-1)*N)
// S: O(N)
func merge(intervals [][]int) [][]int {

	var found = true

	for found {
		found = false
		for i := 0; i < len(intervals)-1; i++ {
			currMin, currMax := findMinMax(intervals[i])
			for j := i + 1; j < len(intervals); j++ {
				nextMin, nextMax := findMinMax(intervals[j])
				if min, max, overlap := overlaps(currMin, nextMin, currMax, nextMax); overlap {
					intervals[i] = []int{min, max}
					intervals = append(intervals[:j], intervals[j+1:]...)
					found = true
					break
				}
			}
		}
	}

	return intervals

}

func overlaps(min1, min2, max1, max2 int) (int, int, bool) {
	min, max := min1, max2
	if min1 > min2 {
		min = min2
	}
	if max1 > max2 {
		max = max1
	}
	// reverse for when they don't overlap
	return min, max, max1 >= min2 && max2 >= min1
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
