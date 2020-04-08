package main

type Node struct {
	Val       int
	Neighbors []*Node
}

type Queue struct {
	nodes []*Node
}

func (q *Queue) enqueue(n *Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) dequeue() *Node {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *Queue) empty() bool {
	return len(q.nodes) == 0
}

func cloneGraph(node *Node) *Node {

	var (
		saw  = map[int]struct{}{}
		seen = func(n int) bool {
			_, ok := saw[n]
			return ok
		}
		q     = &Queue{}
		clone = &Node{}
	)

	q.enqueue(node)

	for !q.empty() {
		cn := q.dequeue()
		if !seen(cn.Val) {
			saw[cn.Val] = struct{}{}
			for _, n := range cn.Neighbors {
				q.enqueue(n)
			}
		}
		clone.Val = cn.Val
		clone.Neighbors = append([]*Node{}, cn.Neighbors...)
	}

	return clone

}

func main() {
	// [[2,4],[1,3],[2,4],[1,3]],
}
