package gameoflife

// We can solve this problem in 2 passes.
// On the first pass, we mark cells that change from dead to alive and viceversa.
//
// We use special values for them in order to not condition the next assertions:
//
// - -1 for a cell that was alive and should become dead - will be replaced by a 0
// - 2 for a cell that was dead and should become alive - will be replaced by a 1
//
// After this phase, we iterate once more the board and finalise the results based on the special values of each cell:
//
// We could solve this by allocating a slice of n elements where we save the new value for each cell.
//
// T: O(r*c)
// S: O(1)
func gameOfLife(board [][]int) {

	const (
		aliveToDead = -1
		alive       = 1
		deadToAlive = 2
	)

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {

			var (
				numAlive int
				isAlive  = alive == board[i][j]
			)

			for _, d := range [][]int{{0, -1}, {-1, -1}, {-1, 1}, {-1, 0}, {0, 1}, {1, 0}, {1, 1}, {1, -1}} {
				x, y := i+d[0], j+d[1]
				if x < 0 || x >= len(board) || y < 0 || y >= len(board[i]) {
					continue
				}
				if alive == board[x][y] || aliveToDead == board[x][y] {
					numAlive++
				}
			}

			switch {
			case isAlive && (numAlive < 2 || numAlive > 3):
				board[i][j] = aliveToDead
			case false == isAlive && 3 == numAlive:
				board[i][j] = deadToAlive
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if aliveToDead == board[i][j] {
				board[i][j] = 0
			} else if deadToAlive == board[i][j] {
				board[i][j] = 1
			}
		}
	}
}
