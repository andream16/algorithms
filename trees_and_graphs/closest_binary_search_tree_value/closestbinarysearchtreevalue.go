package main

import "math"

func main() {}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type Queue struct {
    nodes []*TreeNode
}

func (q *Queue) enqueue(n *TreeNode) {
    q.nodes = append(q.nodes, n)
}

func (q *Queue) dequeue() *TreeNode {
    t := q.nodes[0]
    q.nodes = q.nodes[1:]
    return t
}

func (q *Queue) isEmpty() bool {
    return len(q.nodes) == 0
}

func closestValue(root *TreeNode, target float64) int {
    
    minDist  := float64(math.MaxInt64)
    candidate := 0
    
    queue := &Queue{}
    queue.enqueue(root)
    
    for !queue.isEmpty() {
        n := queue.dequeue()
        if diff := math.Abs(float64(n.Val) - target); diff < minDist {
            minDist = diff
            candidate = n.Val
        }
        if n.Left != nil {
            queue.enqueue(n.Left)
        }
        if n.Right != nil {
            queue.enqueue(n.Right)
        }
    }
    
    return candidate
    
}
