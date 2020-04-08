package main

func main() {}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type LinkedList []*TreeNode

func (l *LinkedList) Add(n *TreeNode) {
    *l = append(*l, n)
}

func flatten(root *TreeNode) {
    l := &LinkedList{}
    preOrderTraversal(root, l)
    
    if ln := len(*l); ln <=1 {
        return
    }
    
    for _, n := range *l {
        root.Left = nil
        root.Right = n
        root = root.Right
    }
}

func preOrderTraversal(n *TreeNode, l *LinkedList) {
    if n == nil {
        return
    }
    l.Add(n)
    preOrderTraversal(n.Left, l)
    preOrderTraversal(n.Right, l)
    return
}
