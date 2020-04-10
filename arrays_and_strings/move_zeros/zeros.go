package main

import "fmt"

// Input: [0,1,0,3,12]
// Output: [1,3,12,0,0]

func moveZeroes(nums []int) {
	var (
		zeros = 0
		l     = len(nums)
	)
	for i := 0; i < len(nums); i++ {
		if l-zeros == i {
			return
		}
		if nums[i] == 0 {
			nums = append(append(nums[:i], nums[i+1:]...), 0)
			zeros++
			i--
		}
	}
}

func main() {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
