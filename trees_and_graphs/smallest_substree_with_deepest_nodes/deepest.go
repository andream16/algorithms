package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LeveledNode struct {
	Node   *TreeNode
	Level  int
	Parent *LeveledNode
}

type queue struct {
	nodes []*LeveledNode
}

func (q *queue) enqueue(node *LeveledNode) {
	q.nodes = append(q.nodes, node)
}

func (q *queue) dequeue() *LeveledNode {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *queue) empty() bool {
	return len(q.nodes) == 0
}

// T: BFS O(N)+(|leaves|*H)
// S: O(N) + |leaves|
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	var (
		deepest []*LeveledNode
		q       = &queue{}
		min     = 0
	)

	q.enqueue(&LeveledNode{
		Node:   root,
		Level:  1,
		Parent: &LeveledNode{Node: root},
	})

	for !q.empty() {

		n := q.dequeue()

		if n.Node.Left == nil && n.Node.Right == nil {
			if n.Level >= min {
				if n.Level == min {
					deepest = append(deepest, n)
				} else {
					min = n.Level
					deepest = []*LeveledNode{n}
				}
			}
		}

		if n.Node.Left != nil {
			q.enqueue(&LeveledNode{
				Node:   n.Node.Left,
				Level:  n.Level + 1,
				Parent: n,
			})
		}

		if n.Node.Right != nil {
			q.enqueue(&LeveledNode{
				Node:   n.Node.Right,
				Level:  n.Level + 1,
				Parent: n,
			})
		}
	}

	if len(deepest) > 0 {
		if len(deepest) <= 1 {
			return deepest[0].Node
		}

		var (
			parent = &TreeNode{}
			same   = true
		)

		for {
			next := []*LeveledNode{}
			for i, n := range deepest {
				next = append(next, n.Parent)
				if i == 0 {
					parent = n.Parent.Node
					continue
				}
				if parent.Val != n.Parent.Node.Val {
					same = false
				}
			}
			if same {
				return parent
			}
			deepest = next
			same = true
		}

	}

	return root

}

func main() {

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	fmt.Println(subtreeWithAllDeepest(root).Val)

}
