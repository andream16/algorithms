package main

import "fmt"

type square struct {
	Val int
	Row int
	Col int
}

type queue struct {
	squares []square
}

func (q *queue) enqueue(square square) {
	q.squares = append(q.squares, square)
}

func (q *queue) dequeue() square {
	square := q.squares[0]
	q.squares = q.squares[1:]
	return square
}

func (q *queue) empty() bool {
	return len(q.squares) == 0
}

// Using BFS, we manipulate and add to the queue only squares with same centerValue and that are inside the image.
//
// T: O(E)
// S: O(V)
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	var (
		centerColor  = image[sr][sc]
		nRows, nCols = len(image), len(image[0])
		visited      = map[string]struct{}{}
	)

	q := &queue{}
	q.enqueue(square{
		Val: centerColor,
		Row: sr,
		Col: sc,
	})

	for !q.empty() {

		sq := q.dequeue()

		for _, m := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			currRow, currCol := sq.Row+m[0], sq.Col+m[1]
			if currRow < 0 || currCol < 0 || currRow >= nRows || currCol >= nCols {
				continue
			}
			currName := getName(currRow, currCol)
			_, ok := visited[currName]
			if image[currRow][currCol] == centerColor && !ok {
				visited[currName] = struct{}{}
				q.enqueue(square{
					Val: image[currRow][currCol],
					Row: currRow,
					Col: currCol,
				})
			}
		}

		image[sq.Row][sq.Col] = newColor

	}

	return image

}

func getName(row, col int) string {
	return fmt.Sprintf("%d%d", row, col)
}

func main() {
	// 2 2 2
	// 2 2 0
	// 2 0 1
	res := floodFill(
		[][]int{
			{1, 1, 1},
			{1, 1, 0},
			{1, 0, 1},
		},
		1,
		1,
		2,
	)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Print(fmt.Sprintf(" %d", res[i][j]))
		}
		fmt.Println("")
	}

}
