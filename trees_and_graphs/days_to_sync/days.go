package main

import "fmt"

func daysToSync(m [][]int) int {

	const (
		notSynced = iota
		synced
	)

	var (
		days = 0
		dirty = false
	)

	for !dirty {
		dirty = true
		hm := deepCopy(m)
		for i:=0; i<len(m); i++ {
			for j:=0; j<len(m[i]); j++ {
				if m[i][j] == synced {
					continue
				}
				dirty = false
				for _, mvs := range [][]int{{-1, 0}, {0,1}, {1, 0}, {0, -1}} {
					currRow, currCol := i+mvs[0], j+mvs[1]
					if currRow < 0 || currCol < 0 || currRow >= len(m) || currCol >= len(m[i]) || m[currRow][currCol] == notSynced {
						continue
					}
					hm[i][j] = synced
					break
				}
			}
		}
		if !dirty {
			days++
		}
		m = hm
	}

	return days

}

func deepCopy(m [][]int) [][]int {
	nm := make([][]int, len(m))
	for i:=0; i<len(m); i++ {
		nm[i] = make([]int, len(m[i]))
		for j:=0; j<len(m[i]); j++ {
			nm[i][j] = m[i][j]
		}
	}
	return nm
}

func main() {
	fmt.Println(daysToSync([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	})) // 6
	fmt.Println(daysToSync([][]int{
		{1, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 1},
	})) // 3
	fmt.Println(daysToSync([][]int{
		{1},
	})) // -1
	fmt.Println(daysToSync([][]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	})) // -1
}
