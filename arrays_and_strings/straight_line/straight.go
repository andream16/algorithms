package main

import "fmt"

// T: O(n)
// S: O(1)
func checkStraightLine(coordinates [][]int) bool {
	switch len(coordinates) {
	case 0, 1:
		return false
	case 2:
		return true
	}
	for i := 1; i < len(coordinates)-2; i++ {
		if 0 != crossProduct(coordinates[0], coordinates[i], coordinates[len(coordinates)-1]) {
			return false
		}
	}
	return true
}

func crossProduct(a, b, c []int) int {
	return det(b, c) - det(a, c) + det(a, b)
}

func det(a, b []int) int {
	return (a[0] * b[1]) - (a[1] * b[0])
}

func main() {
	fmt.Println(checkStraightLine([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}})) // true
	fmt.Println(checkStraightLine([][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}})) // false
}
