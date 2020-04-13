package main

import (
	"fmt"
)

func main() {
	fmt.Println(cutWood([]int{5, 9, 7}, 3))       // 5
	fmt.Println(cutWood([]int{5, 9, 7}, 4))       // 4
	fmt.Println(cutWood([]int{232, 124, 456}, 7)) // 115
	fmt.Println(cutWood([]int{1, 2, 3}, 7))       // 0
}

// T: O(log N)
// S: left, right, middle, bLen, K, int
func cutWood(w []int, k int) int {
	if len(w) == 0 || k == 0 {
		return 0
	}

	n, min, ok := sumAndMinIfValid(w, k)
	if !ok {
		return 0
	}

	return cutWoodRec(1, min, 0, k, k, n)
}

func cutWoodRec(left, right, middle, bLen, k, n int) int {
	middle = (left + right) / 2

	if left == right {
		return bLen
	}

	if (n / middle) >= k {
		bLen = middle
		left = middle + 1
	} else {
		right = middle - 1
	}
	return cutWoodRec(left, right, 0, bLen, k, n)
}

func sumAndMinIfValid(arr []int, k int) (int, int, bool) {
	min, sum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < k {
			return 0, 0, false
		}
		if arr[i] < min {
			min = arr[i]
		}
		sum += arr[i]
	}
	return sum, min, true
}
