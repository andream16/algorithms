package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LeveledNode struct {
	Node  *TreeNode
	Level int
}

type queue struct {
	nodes []*LeveledNode
}

func (q *queue) enqueue(node *LeveledNode) {
	q.nodes = append(q.nodes, node)
}

func (q *queue) dequeue() *LeveledNode {
	n := q.nodes[0]
	q.nodes = append(q.nodes[1:])
	return n
}

func (q *queue) empty() bool {
	return len(q.nodes) == 0
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		min = math.MaxInt64
		q   = &queue{}
	)

	q.enqueue(&LeveledNode{Node: root, Level: 1})

	for !q.empty() {
		n := q.dequeue()
		if n.Node.Left == nil && n.Node.Right == nil {
			if n.Level < min {
				min = n.Level
			}
			continue
		}
		if n.Node.Left != nil {
			q.enqueue(&LeveledNode{Node: n.Node.Left, Level: n.Level + 1})
		}
		if n.Node.Right != nil {
			q.enqueue(&LeveledNode{Node: n.Node.Right, Level: n.Level + 1})
		}
	}

	return min
}

func main() {
	n := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(minDepth(n))
}
