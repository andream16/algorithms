package main

import "fmt"

type node struct {
	x             int
	y             int
	name          string
	depth         int
	fromDirection []int
}

type queue struct {
	nodes []*node
}

func (q *queue) enqueue(n *node) {
	q.nodes = append(q.nodes, n)
}

func (q *queue) dequeue() *node {
	n := q.nodes[0]
	q.nodes = q.nodes[1:]
	return n
}

func (q *queue) empty() bool {
	return len(q.nodes) == 0
}

// T: O(n*m) where M are nodes in the same direction.
// S: O(n*m), O(n*8) for the visited map. 8 for directions.
func longestLine(M [][]int) int {

	var (
		q       = &queue{}
		visited = map[string]map[string]struct{}{}
		longest = 0
	)

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[i]); j++ {

			if M[i][j] == 0 {
				continue
			}

			q.enqueue(&node{x: i, y: j, depth: 1})

			for !q.empty() {
				n := q.dequeue()

				if n.depth > longest {
					longest = n.depth
				}

				if dirs, ok := visited[n.name]; ok {
					dirName := getName(n.fromDirection[0], n.fromDirection[1])
					if _, ok := dirs[dirName]; ok {
						continue
					}
					visited[n.name][dirName] = struct{}{}
				} else {
					visited[getName(n.x, n.y)] = map[string]struct{}{}
				}

				if len(n.fromDirection) == 0 {
					for _, m := range [][]int{{0, -1}, {-1, -1}, {-1, 1}, {-1, 0}, {0, 1}, {1, 0}, {1, 1}, {1, -1}} {
						x, y := n.x+m[0], n.y+m[1]
						if x < 0 || x >= len(M) || y < 0 || y >= len(M[0]) || M[x][y] == 0 {
							continue
						}
						q.enqueue(&node{
							x:             x,
							y:             y,
							name:          getName(x, y),
							depth:         n.depth + 1,
							fromDirection: m,
						})
					}
				} else {
					x, y := n.x+n.fromDirection[0], n.y+n.fromDirection[1]
					if x < 0 || x >= len(M) || y < 0 || y >= len(M[0]) || M[x][y] == 0 {
						continue
					}
					q.enqueue(&node{
						x:             x,
						y:             y,
						name:          getName(x, y),
						depth:         n.depth + 1,
						fromDirection: n.fromDirection,
					})
				}

			}

		}
	}

	return longest

}

func getName(x, y int) string {
	return fmt.Sprintf("%d%d", x, y)
}

func main() {
	fmt.Println(longestLine([][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	})) // 3
	fmt.Println(longestLine([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	})) // 3
}
