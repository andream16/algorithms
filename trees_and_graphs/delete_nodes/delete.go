package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// T: O(D * N * N)
// S: O(T)
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	if root == nil || len(to_delete) == 0 {
		return []*TreeNode{root}
	}

	trees := []*TreeNode{root}

	for _, n := range to_delete {
		for i, t := range trees {
			if toBeDeleted := dfs(t, n); toBeDeleted != nil {
				if n == t.Val && toBeDeleted.Left == nil && toBeDeleted.Right == nil {
					trees = append(append(trees[:i], trees[i+1:]...))
					break
				}
				if toBeDeleted.Left != nil {
					trees = append(trees, toBeDeleted.Left)
				}
				if toBeDeleted.Right != nil {
					trees = append(trees, toBeDeleted.Right)
				}
				if n != t.Val {
					ogNode := TreeNode{}
					deepCopy(t, &ogNode, n)
					trees = append(append(trees[:i], trees[i+1:]...), &ogNode)
					break
				}
				trees = append(append(trees[:i], trees[i+1:]...))
				break
			}
		}
	}

	return trees
}

func dfs(node *TreeNode, target int) *TreeNode {
	if node.Val == target {
		return node
	}
	if node.Right == nil && node.Left == nil {
		return nil
	}
	if node.Right != nil {
		if f := dfs(node.Right, target); f != nil {
			return f
		}
	}
	if node.Left != nil {
		if f := dfs(node.Left, target); f != nil {
			return f
		}
	}
	return nil
}

func deepCopy(node, newNode *TreeNode, target int) {
	newNode.Val = node.Val
	if node.Right != nil && node.Right.Val != target {
		newNode.Right = &TreeNode{}
		deepCopy(node.Right, newNode.Right, target)
	}
	if node.Left != nil && node.Left.Val != target {
		newNode.Left = &TreeNode{}
		deepCopy(node.Left, newNode.Left, target)
	}
}

func main() {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	delNodes(node, []int{2, 1})
}
