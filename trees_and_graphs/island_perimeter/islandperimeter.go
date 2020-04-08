package main

import (
	"fmt"
)

func main()  {
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

type land struct {
	x int
	y int
	neighbours []land
}

func fillNeighbours(island []land, maxLen int) {
	for i:=0; i<len(island); i++ {
		for j:=0; j<len(island); j++ {
			if isNextTo(island[i], island[j], maxLen) {
				island[i].neighbours = append(island[i].neighbours, island[j])
			}
		}
	}
}

func isNextTo(l1, l2 land, maxLen int) bool {
	if l1.x == l2.x && l1.y	== l2.y {
		return false
	}
	if l1.x > 0 && l1.y > 0 && l1.x < maxLen && l1.y < maxLen {
		return (l1.y - 1 == l2.y && l1.x == l2.x) || (l1.y + 1 == l2.y && l1.x == l2.x) || (l1.x - 1 == l2.x && l1.y == l2.y) || (l1.x + 1 == l2.x && l1.y == l2.y)
	}
	if l1.x == 0 {
		if l1.y == 0 {
			return (l1.y + 1 == l2.y && l1.x == l2.x) || (l1.x + 1 == l2.x && l1.y == l2.y)
		}
		if l1.y == maxLen {
			return (l1.y - 1 == l2.y && l1.x == l2.x) || (l1.x + 1 == l2.x && l1.y == l2.y)
		}
		return (l1.y - 1 == l2.y && l1.x == l2.x) || (l1.y + 1 == l2.y && l1.x == l2.x) || (l1.x + 1 == l2.x && l1.y == l2.y)
	}
	if l1.x == maxLen {
		if l1.y == 0 {
			return (l1.y + 1 == l2.y && l1.x == l2.x) || (l1.x - 1 == l2.x && l1.y == l2.y)
		}
		if l1.y == maxLen {
			return (l1.y - 1 == l2.y && l1.x == l2.x) || (l1.x - 1 == l2.x && l1.y == l2.y)
		}
		return (l1.y - 1 == l2.y && l1.x == l2.x) || (l1.y + 1 == l2.y && l1.x == l2.x) || (l1.x - 1 == l2.x && l1.y == l2.y)
	}
	if l1.y == 0 {
		return (l1.y + 1 == l2.y && l1.x == l2.x) || (l1.x + 1 == l2.x && l1.y == l2.y) || (l1.x - 1 == l2.x && l1.y == l2.y)
	}
	return (l1.y - 1 == l2.y && l1.x == l2.x) || (l1.x + 1 == l2.x && l1.y == l2.y) || (l1.x - 1 == l2.x && l1.y == l2.y)
}

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	var (
		island = []land{}
		perimeter = 0
	)

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}
			island = append(island, land{
				x:          i,
				y:          j,
			})
		}
	}

	fillNeighbours(island, len(grid[0])-1)

	for _, l := range island {
		switch len(l.neighbours) {
		case 0:
			perimeter += 4
		case 1:
			perimeter += 3
		case 2:
			perimeter += 2
		}
	}

	return perimeter
}
