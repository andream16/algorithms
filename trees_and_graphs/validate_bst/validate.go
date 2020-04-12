package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(n *TreeNode, min, max int) bool {
	if n == nil {
		return true
	}
	if n.Val <= min || n.Val >= max {
		return false
	}
	if n.Left != nil && n.Left.Val >= n.Val {
		return false
	}
	if n.Right != nil && n.Right.Val <= n.Val {
		return false
	}
	return helper(n.Left, min, n.Val) && helper(n.Right, n.Val, max)
}

func main() {}
