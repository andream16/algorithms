package maxprofit

import (
	"math"
)

// To solve this problem, we need the concepts of minimum price and maximum difference.
// We iterate the prices and every time we find a price that is lower than the minimum encountered, we update the latter.
// The most profitable day, is the day where the difference between the current price and the minimum is the greatest.
//
// T: O(n)
// S: O(n)
func maxProfit(prices []int) int {
	var (
		max = 0
		min = math.MaxInt64
	)

	for _, price := range prices {
		if price < min {
			min = price
			continue
		}
		if diff := price - min; diff > max {
			max = diff
		}
	}

	return max
}
