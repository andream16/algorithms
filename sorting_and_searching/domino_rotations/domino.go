package main

import (
	"fmt"
)

func minDominoRotations(A []int, B []int) int {
	r := rotations(A[0], A, B)
	if -1 != r || A[0] == B[0] {
		return r
	}
	return rotations(B[0], A, B)
}

func rotations(n int, A, B []int) int {
	rotationsA, rotationsB := 0, 0
	for i := 0; i < len(A); i++ {
		if n != A[i] && n != B[i] {
			return -1
		} else if n != A[i] {
			rotationsA++
		} else if n != B[i] {
			rotationsB++
		}
	}
	return min(rotationsA, rotationsB)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2})) // 2
	fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))       // -1
}
