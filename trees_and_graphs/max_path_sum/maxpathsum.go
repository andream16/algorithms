package maxpathsum

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// To solve this problem we can use Depth Search First and a concept of maximum gain.
// When we visit a node, we check the current gain of it by summing up its value and the gain of the various nodes
// in its subtree. Every time we do so, we check if the current gain is greater than the maximum gain. If so, we update
// the latter.
//
// T: O(N)
// S: O(N)
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return root.Val
	}
	maxSum := math.MinInt32
	findMax(root, &maxSum)
	return maxSum
}

func findMax(n *TreeNode, maxSum *int) int {
	if n == nil {
		return 0
	}

	leftGain, rightGain := 0, 0
	if n.Left != nil {
		leftGain = max(findMax(n.Left, maxSum), 0)
	}
	if n.Right != nil {
		rightGain = max(findMax(n.Right, maxSum), 0)
	}

	currGain := n.Val + leftGain + rightGain
	*maxSum = max(*maxSum, currGain)

	return n.Val + max(leftGain, rightGain)
}

func max(n1, n2 int) int {
	if n1 >= n2 {
		return n1
	}
	return n2
}
