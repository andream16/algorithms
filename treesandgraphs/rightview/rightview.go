/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type LeveledNode struct {
    Node *TreeNode
    Level int
}

type queue struct {
    nodes []*LeveledNode
}

func (q *queue) enqueue(ln *LeveledNode) {
    if len(q.nodes) == 0 {
        q.nodes = []*LeveledNode{}
    }
    q.nodes = append(q.nodes, ln)
}

func (q *queue) dequeue() *LeveledNode {
    n := q.nodes[0]
    q.nodes = q.nodes[1:]
    return n
}

func (q *queue) empty() bool {
    return len(q.nodes) == 0
}

func newLeveledNode(n *TreeNode, lvl int) *LeveledNode {
    return &LeveledNode{
        Node: n,
        Level: lvl,
    }
}

func seen(n int, h map[int]struct{}) bool {
    _, ok := h[n]
    return ok
}

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return make([]int, 0)
    }
    
    vals := []int{}
    bucket := map[int]struct{}{}
    q := &queue{}
    q.enqueue(newLeveledNode(root, 0))
    
    for !q.empty() {
        ln := q.dequeue()
        if !seen(ln.Level, bucket) {
           vals = append(vals, ln.Node.Val)
            bucket[ln.Level] = struct{}{}
        }
        
        if ln.Node.Right != nil {
            q.enqueue(newLeveledNode(ln.Node.Right, ln.Level+1))
        }
        if ln.Node.Left != nil {
            q.enqueue(newLeveledNode(ln.Node.Left, ln.Level+1))
        }
    }
    
    return vals
}
