package main

import "fmt"

// T: O(n)
// S: O(n)

// There's only one exception in case the array has length 3.
// In all the other cases, until we find the peek, point where A[i-1] < A[i] > A[i+1], we check for each element
// to be greater than the previous A[i] > A[i-1]. When instead we find the peek, we need to check that each element
// is smaller than the previous one A[i] < A[i-1].
func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}

	var peekIndex = -1

	for i := 1; i <= len(A)-1; i++ {
		if i < len(A)-1 && A[i-1] < A[i] && A[i] > A[i+1] {
			if peekIndex > 0 {
				return false
			}
			peekIndex = i
			continue
		}
		if (peekIndex < 0 && A[i-1] >= A[i]) || (peekIndex > 0 && A[i-1] <= A[i]) {
			return false
		}
	}

	if peekIndex == -1 {
		return false
	}

	return true
}

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))                                // false
	fmt.Println(validMountainArray([]int{3, 5, 5}))                             // false
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))                          // true
	fmt.Println(validMountainArray([]int{0, 2, 3, 4, 5, 2, 1, 0}))              // true
	fmt.Println(validMountainArray([]int{0, 2, 3, 3, 5, 2, 1, 0}))              // false
	fmt.Println(validMountainArray([]int{1, 3, 2}))                             // true
	fmt.Println(validMountainArray([]int{0, 1, 2, 1, 2}))                       // false
	fmt.Println(validMountainArray([]int{3, 7, 20, 14, 15, 14, 10, 8, 2, 1}))   // false
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4, 8, 9, 10, 11, 12, 11})) // true
}
