package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// TODO Work on this
func countSmaller(nums []int) []int {
	sort.Ints(nums)
	var res []int
	_ = buildBST(nums, &res)
	return res
}

func buildBST(arr []int, res *[]int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr)/2
	*res = append(*res, mid)
	return &TreeNode{
		Val:   arr[mid],
		Left:  buildBST(arr[:mid], res),
		Right: buildBST(arr[mid+1:], res),
	}
}

func main() {
	fmt.Println(countSmaller([]int{5, 2, 6, 1})) // [2, 1, 1, 0]
}
