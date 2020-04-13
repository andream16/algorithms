package main

import "fmt"

type NumMatrix struct {
	matrix [][]int
}

// T: constructor O(R*C) + sumRegion O(SR*SC)
// S: r*c, 1 for sum
func Constructor(matrix [][]int) NumMatrix {
	m := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		m[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[0]); j++ {
			m[i][j] = matrix[i][j]
		}
	}
	return NumMatrix{matrix: m}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sum += this.matrix[i][j]
		}
	}
	return sum
}

// Given matrix = [
//  [3, 0, 1, 4, 2],
//  [5, 6, 3, 2, 1],
//  [1, 2, 0, 1, 5],
//  [4, 1, 0, 1, 7],
//  [1, 0, 3, 0, 5]
//]
func main() {
	m := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	nm := Constructor(m)
	// sumRegion(2, 1, 4, 3) -> 8
	//sumRegion(1, 1, 2, 2) -> 11
	//sumRegion(1, 2, 2, 4) -> 12
	fmt.Println(nm.SumRegion(1, 2, 2, 4))
}
