package main

import "fmt"

// We can solve this using binary search. We use two pointers, high = len(A)-1 and low = 0.
// If the current value, given by mid = (low+high)/2, is smaller than the previous one, it means
// that the peak is on the right side of it, therefore we look for it by assigning low = mid+1.
// On the other hand, if this is bigger, it means that the peak is on the left side, therefore,
// we assign high = mid. We stop when the sto pointers meet.
//
// T: O(log n)
// S: O(1)
func peakIndexInMountainArray(A []int) int {
	return binarySearchPeak(0, len(A)-1, A)
}

func binarySearchPeak(lowIdx, highIdx int, arr []int) int {
	if lowIdx == highIdx {
		return lowIdx
	}
	middleIdx := (lowIdx+highIdx) / 2
	if arr[middleIdx] < arr[middleIdx+1] {
		return binarySearchPeak(middleIdx+1, highIdx, arr)
	}
	return binarySearchPeak(lowIdx, middleIdx, arr)
}

// T: O(n)
// S: O(1)
func peakIndexInMountainArrayBasic(A []int) int {
	for i:=1; i<len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			return i
		}
	}
	return 0
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0,1,0})) // 1
	fmt.Println(peakIndexInMountainArray([]int{0,2,1,0})) // 1
	fmt.Println(peakIndexInMountainArray([]int{0,2,3,4, 5, 6, 2, 1})) // 5
}
