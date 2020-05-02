package main

import "fmt"

// T: O(n)
// S: O(1)
func duplicateZeros(arr []int) {
	duplicateZerosRec(arr, 0)
}

func duplicateZerosRec(arr []int, nextIdx int) {
	if nextIdx > len(arr)-2 {
		return
	}
	if arr[nextIdx] != 0 {
		duplicateZerosRec(arr, nextIdx+1)
		return
	}
	duplicateZerosRec(append(append(arr[:nextIdx], 0), arr[nextIdx:len(arr)-1]...), nextIdx+2)
}

func main() {
	arr1 := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr1)
	fmt.Println(arr1)
	arr2 := []int{1, 2, 3}
	duplicateZeros(arr2)
	fmt.Println(arr2)
}
