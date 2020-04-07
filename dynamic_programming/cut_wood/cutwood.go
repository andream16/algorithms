package main

import (
	"fmt"
)

func main() {
	fmt.Println(cutWood([]int{5, 9, 7}, 4))
}

func cutWood(w []int, k int) int {
	if len(w) == 0 || k == 0 {
		return 0
	}
	
	totPieces, min := sumAndMin(w)
	bLen, left, right, middle := 0, 1, min, 0
	
	for right != left {
		middle = (left+right) / 2
		if totPieces / middle >= k {
			left = middle
		} else {
			right= middle
		}
		bLen = middle
	}
	
	return bLen
	
}

func sumAndMin(arr []int) (int, int) {
	min, sum := arr[0], arr[0]
	for i:=1; i<len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		sum += arr[i]
	}
	return sum, min
}
