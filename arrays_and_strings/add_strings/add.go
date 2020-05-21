package add

import "fmt"

// Keep adding digit by digits starting from right to left. Keep in mind the carry.
//
// T: O(n) where n is the longest.
// S: O(1)
func addStrings(num1 string, num2 string) string {

	var (
		sum        string
		carry      int
		idx1, idx2 = len(num1) - 1, len(num2) - 1
	)

	for idx1 >= 0 || idx2 >= 0 {

		var curr1, curr2, currSum = 0, 0, 0

		if idx1 >= 0 {
			curr1 = int(num1[idx1] - '0')
			idx1 -= 1
		}
		if idx2 >= 0 {
			curr2 = int(num2[idx2] - '0')
			idx2 -= 1
		}
		currSum = curr1 + curr2 + carry
		carry = currSum / 10
		currSum = currSum % 10

		sum = fmt.Sprintf("%d", currSum) + sum

	}

	if carry > 0 {
		sum = fmt.Sprintf("%d", carry) + sum
	}

	return sum
}
