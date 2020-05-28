package diameter

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
}

// The diameter of a binary tree is the length of the longest path between two arbitrary nodes.
// It's important to notice that this path may not pass through the root node.
// To find the diameter, we to calculate recursively the depth of each left and right sub-trees and consider what's
// greater between their diameter (depthR + depthL + 1) and the current greatest one. +1 is used to increment the level.
// For the recursive call, we use max(left, right) + 1 because we have to pick only one path per tree.
//
// T: O(E)
// S: O(E)
func diameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	dfs(root, &diameter)
	return diameter
}

func dfs(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}
	var left, right int
	if node.Left != nil {
		left = dfs(node.Left, diameter)
	}
	if node.Right != nil {
		right = dfs(node.Right, diameter)
	}
	*diameter = max(*diameter, left+right)
	return max(left, right) + 1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
