package multiply_strings

import (
	"strconv"
	"strings"
)

// We could solve this naively by adding up all inflight results for each digit.
// This method, instead, is more clean even if a bit more complicated.
// We use a results array which has dimensions len(n1) + len(n2) + 1 as this is the maximum amount
// of digits that a product calculated from n1 and n1 can have.
// Using the indexes, we add the current values of each multiplication at the right location.
// We use the same principle to take care of the carry.
// Reversing the strings & reversing them back makes it easier to work with.
//
// T: O(m*n)
// S: O(len(m)+len(n))
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var (
		// We can have max len(num1) + len(num2) + 1 digit long.
		res            = make([]int, len(num1)+len(num2)+1)
		idxToBeRemoved = 0
		ans            string
	)

	// Reversing simplifies the whole process - otherwise we'd have to start from the
	// last digits of each number.
	num1, num2 = reverseStr(num1), reverseStr(num2)

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			// Summing up at location i+j.
			res[i+j] += int(num1[i]-'0') * int(num2[j]-'0')
			// Registering the carry to the next location.
			res[i+j+1] += res[i+j] / 10
			// What's left without the carry will be in the current location.
			res[i+j] = res[i+j] % 10
		}
	}

	// Removing trailing 0s.
	for i := len(res) - 1; i > 0; i-- {
		if res[i] > 0 {
			break
		}
		idxToBeRemoved = i
	}

	res = res[:idxToBeRemoved]

	for i := 0; i < len(res); i++ {
		ans = strconv.Itoa(res[i]) + ans
	}

	return ans
}

func reverseStr(s string) string {
	var rs strings.Builder

	for i := len(s) - 1; i >= 0; i-- {
		rs.WriteByte(s[i])
	}

	return rs.String()
}
