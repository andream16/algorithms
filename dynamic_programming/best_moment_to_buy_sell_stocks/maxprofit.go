package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	var (
		max = 0
		min = math.MaxInt64
	)

	for _, p := range prices {
		if p < min {
			min = p
			continue
		}
		if diff := p - min; diff > max {
			max = diff
		}
	}

	return max
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))    // 5
	fmt.Println(maxProfit([]int{7, 6, 5, 4, 3, 2, 1})) // 0
}
