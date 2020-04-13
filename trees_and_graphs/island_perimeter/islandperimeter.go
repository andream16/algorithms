package main

import (
	"fmt"
)

func main() {
	//m := [][]int{
	//	{0, 1, 0, 0},
	//	{1, 1, 1, 0},
	//	{0, 1, 0, 0},
	//	{1, 1, 0, 0},
	//}
	m := [][]int{
		{0, 1, 1},
		{0, 0, 0},
	}
	fmt.Println(islandPerimeter(m))
}

// T: O(N*4) -> O(N)
// S: M*N, one var for Perimeter
func islandPerimeter(grid [][]int) int {

	var (
		perimeter = 0
		x         = 0
		y         = 0
		lx        = len(grid)
		ly        = len(grid[0])
	)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			perimeter += 4
			for _, diff := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				x, y = i+diff[0], j+diff[1]
				if x < 0 || x >= lx || y < 0 || y >= ly {
					continue
				}
				if grid[x][y] == 1 {
					perimeter--
				}
			}
		}
	}

	return perimeter
}
