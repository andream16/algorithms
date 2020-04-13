package main

import (
	"fmt"
	"math"
)

type island struct {
	x int
	y int
}

func numIslands(grid [][]byte) int {
	var islands [][]island

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '0' {
				continue
			}
			addToIslands(i, j, &islands)
		}
	}

	return len(islands)
}

// T: BFS O(N*|islands*|land||) -> O(N*I*L)
// S: |islands*lands|
func addToIslands(x, y int, islands *[][]island) {
	var boundaries []int
	for idx, isl := range *islands {
		for _, i := range isl {
			dx, dy := int(math.Abs(float64(i.x-x))), int(math.Abs(float64(i.y-y)))
			if dx == 0 && dy <= 1 || dx <= 1 && dy == 0 {
				boundaries = append(boundaries, idx)
				(*islands)[idx] = append((*islands)[idx], island{x: x, y: y})
				break
			}
		}
	}
	if len(boundaries) > 0 {
		if len(boundaries) > 1 {
			var resIsl []island
			for _, b := range boundaries {
				for _, isl := range (*islands)[b] {
					resIsl = append(resIsl, isl)
				}
				*(islands) = append(append([][]island{}, (*islands)[b]), (*islands)[b+1:]...)
			}
			(*islands)[boundaries[0]] = resIsl
			return
		}
		return
	}
	*islands = append(*islands, []island{{x: x, y: y}})
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}))
}
