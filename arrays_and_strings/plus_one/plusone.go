package plusone

// Assuming that every position has an integer n and 0 <= n <= 9,
// if we find n[i] + 1 == 10, we need to set the current n[i] = 0, and keep adding 1 to the next
// positions until n[i+1]+1 <= 9. We need to add a new position at the head of the slice if the last one is a 9.
//
// T: O(n)
// S O(1)
func plusOne(digits []int) []int {

	needsCarry := true

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] <= 9 {
			needsCarry = false
			break
		}
		digits[i] = 0
	}

	if needsCarry {
		return append([]int{1}, digits...)
	}

	return digits

}
