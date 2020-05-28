package island

import (
	"fmt"
)

type cell struct {
	x int
	y int
}

// To solve this problem, we can hold the visited lands into a set and use BFS to form the islands.
// If a cell is a land and we haven't visited it, use BFS by adding it into a queue and keep adding to the queue all the
// adjacent lands. Every time we do this, it means that we found a new island.
//
// T: O(r x c) where r is the number of rows and c is the number of columns.
// S: O(r x c) as in the worst case the queue can contain all the elements.
func numIslands(grid [][]byte) int {
	var (
		visited  = map[string]struct{}{}
		nIslands int
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			cellName := toCell(i, j)
			if _, ok := visited[cellName]; ok || grid[i][j] == '0' {
				continue
			}
			bfs(i, j, grid, visited)
			nIslands++
		}
	}

	return nIslands
}

func bfs(i, j int, grid [][]byte, visited map[string]struct{}) {
	var queue []cell
	queue = append(queue, cell{x: i, y: j})
	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]
		for _, d := range [][]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}} {
			x, y := c.x+d[0], c.y+d[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[i]) {
				cellName := toCell(x, y)
				if _, ok := visited[cellName]; ok || grid[x][y] == '0' {
					continue
				}
				visited[cellName] = struct{}{}
				queue = append(queue, cell{x: x, y: y})
			}
		}
	}
}

func toCell(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}
