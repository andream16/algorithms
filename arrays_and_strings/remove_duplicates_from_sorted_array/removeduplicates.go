package main

import "fmt"

func main()  {
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}

func removeDuplicates(nums []int) int {
	saw := map[int]struct{}{}
	for i:=0; i<len(nums); i++ {
		if _, ok := saw[nums[i]]; ok {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			continue
		}
		saw[nums[i]] = struct{}{}
	}
	return len(nums)
}
