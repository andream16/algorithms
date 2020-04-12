package main

import "math"

func divide(dividend int, divisor int) int {

	if dividend <= math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	dividend, ok1 := abs(dividend)
	divisor, ok2 := abs(divisor)

	currentDivisor, partialQuotient, finalQuotient := divisor, 1, 0

	for dividend-currentDivisor >= 0 {
		finalQuotient = finalQuotient + partialQuotient
		dividend = dividend - currentDivisor

		partialQuotient, currentDivisor = partialQuotient<<1, currentDivisor<<1
		if dividend-currentDivisor < 0 {
			partialQuotient, currentDivisor = 1, divisor
		}
	}

	if (ok1 && !ok2) || (!ok1 && ok2) {
		return -finalQuotient
	}
	return finalQuotient
}

func abs(n int) (int, bool) {
	mask := n >> 32
	return (n + mask) ^ mask, n < 0
}

func main() {}
