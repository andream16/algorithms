package main

func main() {}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type histNode struct {
    node *TreeNode
    backtrack []int
}

func newHistNode(n *TreeNode, backtrack []int) *histNode {
    return &histNode{
        node: n,
        backtrack: append(append([]int{}, backtrack...), n.Val),
    }
}

type queue struct {
    nodes []*histNode
}

func (q *queue) enqueue(n *histNode) {
    q.nodes = append(q.nodes, n)
}

func (q *queue) dequeue() *histNode {
    n := q.nodes[0]
    q.nodes = q.nodes[1:]
    return n
}

func (q *queue) empty() bool {
    return len(q.nodes) == 0
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if p == nil && q == nil {
        return root
    }
    
    var candidate1 *histNode = nil
    var candidate2 *histNode = nil
    que := &queue{}
    que.enqueue(newHistNode(root, nil))
    
    for candidate1 == nil || candidate2 == nil {
        r := que.dequeue()
                
        if r.node.Val == p.Val {
            candidate1 = r
        } else if r.node.Val == q.Val {
            candidate2 = r
        }
        
        if r.node.Left != nil {
            que.enqueue(newHistNode(r.node.Left, r.backtrack))
        }
        if r.node.Right != nil {
            que.enqueue(newHistNode(r.node.Right, r.backtrack))
        }
    }
    
    for i:=len(candidate1.backtrack)-1; i>=0; i-- {
        for j:=len(candidate2.backtrack)-1; j>=0; j-- {
            if candidate1.backtrack[i] == candidate2.backtrack[j] {
                return &TreeNode{
                    Val: candidate2.backtrack[j],
                }
            }
        }
    }

    return root
}
