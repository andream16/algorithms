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

// T: BFS O(N*|neighbours|)
// S: N for map, M for subnodes in queue
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	var (
		visited        = map[*Node]*Node{}
		hasBeenVisited = func(n *Node) bool {
			_, ok := visited[n]
			return ok
		}
		q = &Queue{}
	)

	q.enqueue(node)

	for !q.empty() {
		cn := q.dequeue()
		for _, n := range cn.Neighbors {
			if !hasBeenVisited(n) {
				visited[n] = &Node{
					Val:       n.Val,
					Neighbors: nil,
				}
				q.enqueue(n)
			}
			if !hasBeenVisited(cn) {
				visited[cn] = &Node{
					Val: cn.Val,
				}
			}
			visited[cn].Neighbors = append(visited[cn].Neighbors, visited[n])
		}
	}

	clone := visited[node]
	if clone == nil || clone == (&Node{}) {
		r := &Node{}
		*r = *node
		return r
	}

	return clone

}

func main() {
	n := &Node{
		Val: 1,
		Neighbors: []*Node{
			{
				Val: 2,
				Neighbors: []*Node{
					{
						Val: 1,
					},
					{
						Val: 3,
						Neighbors: []*Node{
							{
								Val: 2,
							},
							{
								Val: 4,
								Neighbors: []*Node{
									{
										Val: 1,
									},
									{
										Val: 3,
									},
								},
							},
						},
					},
				},
			},
			{
				Val: 4,
				Neighbors: []*Node{
					{
						Val: 1,
					},
					{
						Val: 3,
					},
				},
			},
		},
	}
	// [[2,4],[1,3],[2,4],[1,3]],
	cloneGraph(n)
}
