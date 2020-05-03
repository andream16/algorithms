package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// T: O(n)
// S: O(n)
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root, 0)
}

func dfs(node *TreeNode, count int) int {
	if node != nil {
		count++
	}
	if node.Left != nil {
		count = dfs(node.Left, count)
	}
	if node.Right != nil {
		count = dfs(node.Right, count)
	}
	return count
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
		},
	}
	fmt.Println(countNodes(node)) // 6
}
