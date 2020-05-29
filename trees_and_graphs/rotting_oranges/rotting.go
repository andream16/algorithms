package rotting

type cell struct {
	x int
	y int
}

// This problem can be solved using a tweaked version of BFS.
// We add to the queue all the rotten oranges and an helper that lets us understand if a cycle (minute) has completed.
// We also track the number of fresh oranges to be used to check if have unreachable ones and return -1 accordingly.
// When we pop from the queue, we check if it's the helper cell. If so, it means that one minute has elapsed.
// Otherwise, we check current's cell neighbours and since the current one is rotten, we add to the queue the new
// potential rotten ones while also updating them in the original grid. When we do so, we also decrease the number of
// fresh oranges. When we are done, we check if the number of fresh oranges is 0 or not. If It's zero, it means that
// there were no unreachable oranges and we can therefore return the elapsed number of minutes. We return -1 otherwise.
//
// T: O(r x c)
// S: O(r x c) because in the worst case all cells can go into the queue
func orangesRotting(grid [][]int) int {

	const (
		_ = iota
		freshOrange
		rottenOrange
	)

	var (
		queue                  []cell
		minutes, nFreshOranges = -1, 0
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case freshOrange:
				nFreshOranges += 1
			case rottenOrange:
				queue = append(queue, cell{x: i, y: j})
			}
		}
	}

	queue = append(queue, cell{x: -1, y: -1})

	for len(queue) > 0 {

		c := queue[0]
		queue = queue[1:]
		if -1 == c.x && -1 == c.y {
			minutes++
			if len(queue) > 0 {
				queue = append(queue, cell{x: -1, y: -1})
			}
			continue
		}

		for _, d := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			x, y := c.x+d[0], c.y+d[1]
			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0]) {
				if freshOrange == grid[x][y] {
					grid[x][y] = rottenOrange
					nFreshOranges -= 1
					queue = append(queue, cell{x: x, y: y})
				}
			}
		}

	}

	if 0 == nFreshOranges {
		return minutes
	}

	return -1

}
