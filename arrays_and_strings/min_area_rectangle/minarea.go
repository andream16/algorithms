package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x int
	y int
}

type ByValue []Point

func (b ByValue) Len() int { return len(b) }
func (b ByValue) Less(i, j int) bool {
	return b[i].x < b[j].x
}
func (b ByValue) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

// TODO CHECK AGAIN
func minAreaRect(points [][]int) int {
	if len(points) < 4 {
		return 0
	}

	var (
		pts     = make([]Point, len(points))
		minArea = 0
	)

	for i, pt := range points {
		pts[i] = Point{x: pt[0], y: pt[1]}
	}

	sort.Sort(ByValue(pts))

	for i := 0; i < len(pts); i++ {
		ctr, width, height := 1, 0, 0
		for j := 0; j < len(pts); j++ {
			switch ctr {
			case 1:
				if pts[i].y == pts[j].y && abs(pts[i].x-pts[j].x) > 0 {
					width = abs(pts[i].x - pts[j].x)
					ctr++
					i, j = j, 0
				}
			case 2:
				if pts[i].x == pts[j].x && abs(pts[i].y-pts[j].y) > 0 {
					height = abs(pts[i].y - pts[j].y)
					ctr++
					i, j = j, 0
				}
			case 3:
				if pts[i].y == pts[j].y && abs(pts[i].x-pts[j].x) == width {
					ctr++
					i, j = j, 0
					continue
				}
			case 4:
				if pts[i].x == pts[j].x && abs(pts[i].y-pts[j].y) == height {
					if area := width * height; area < minArea {
						minArea = area
					}
				}
			}
		}
	}

	return minArea
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {2, 2}})) // 4
	//fmt.Println(minAreaRect([][]int{{1, 1}, {1, 3}, {3, 1}, {3, 3}, {4, 1}, {4, 3}})) // 2
}
