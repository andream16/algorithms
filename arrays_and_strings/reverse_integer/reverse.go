package reverseinteger

import (
	"math"
	"strconv"
)

// Check digit by digits if we overflow or not.
// Complexity is O(log n) because we don't need to process all the numbers
//
// T: O(log n)
// S: O(1)
func reverse(x int) int {
	ans := 0

	for x != 0 {
		last := x % 10
		x /= 10
		next := ans*10 + last
		if (next-last)/10 != ans || next > math.MaxInt32 || next < math.MinInt32 {
			return 0
		}
		ans = next
	}

	return ans
}

// We can cast the integer to string and then process character by character.
// There are some edge cases for overflow, negative & trailing 0s numbers.
//
// T: O(n)
// S: O(n)
func reverseBasic(x int) int {
	s := strconv.Itoa(x)

	var (
		res        string
		isNegative = s[0] == '-'
	)

	for _, r := range s {
		if r == '-' {
			continue
		}
		res = string(r) + res
	}

	if len(res) > 1 {
		for {
			if res[0] != '0' {
				break
			}
			res = res[1:]
		}
	}

	n, err := strconv.Atoi(res)
	if err != nil || n > math.MaxInt32 {
		return 0
	}

	if isNegative {
		if -n < math.MinInt32 {
			return 0
		}
		return -n
	}

	return n

}
