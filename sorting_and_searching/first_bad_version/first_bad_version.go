package main

import "fmt"

// T: O(log n)
// S: O(1)
func firstBadVersion(n int) int {
	return helper(0, n)
}

func helper(l, r int) int {
	if l == r {
		return 0
	}
	mid := (l + r) / 2
	if isBadVersion(mid) {
		if !isBadVersion(mid - 1) {
			return mid
		}
		return helper(l, mid-1)
	}
	if isBadVersion(mid + 1) {
		return mid + 1
	}
	return helper(mid+1, r)
}

func isBadVersion(version int) bool {
	return false
}

func main() {
	fmt.Println(firstBadVersion(5))
}
