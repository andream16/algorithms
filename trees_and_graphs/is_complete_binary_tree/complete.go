package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queue struct {
	nodes []*TreeNode
}

func (q *queue) enqueue(n *TreeNode) {
	q.nodes = append(q.nodes, n)
}

func (q *queue) dequeue() *TreeNode {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *queue) empty() bool {
	return len(q.nodes) == 0
}

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var (
		q          = &queue{}
		emptyRight = false
	)

	q.enqueue(root)

	for !q.empty() {
		n := q.dequeue()

		if n.Left == nil && n.Right != nil {
			return false
		}

		if n.Left != nil {
			// As soon as we find an empty Right child we can stop.
			if emptyRight {
				return false
			}
			q.enqueue(n.Left)
		}
		if n.Right != nil {
			q.enqueue(n.Right)
		} else {
			emptyRight = true
		}
	}

	return true

}

func main() {
	n := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 6,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(isCompleteTree(n)) // true
}
