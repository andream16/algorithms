package main

import (
	"fmt"
	"math/rand"
)

type matrix struct {
	nRows    int
	nColumns int
	values   [][]int
}

func (m matrix) Print() {
	for i := 0; i < m.nRows; i++ {
		for j := 0; j < m.nColumns; j++ {
			fmt.Printf("%d ", m.values[i][j])
		}
		fmt.Println("")
	}
}

func newMatrix(nRows, nColumns int) matrix {
	m := matrix{
		nRows:    nRows,
		nColumns: nColumns,
		values:   make([][]int, nRows),
	}
	for i := 0; i < nRows; i++ {
		m.values[i] = make([]int, nColumns)
		for j := 0; j < nColumns; j++ {
			m.values[i][j] = rand.Intn(2)
		}
	}
	return m
}

func (m matrix) canBySynced(r, c int) bool {
	if m.values[r][c] == 1 {
		return false
	}
	switch r {
	case 0:
		switch c {
		case 0:
			return m.values[r+1][c] == 1 || m.values[r][c+1] == 1
		case m.nColumns - 1:
			return m.values[r][c-1] == 1 || m.values[r+1][c] == 1
		default:
			return m.values[r][c-1] == 1 || m.values[r][c+1] == 1 || m.values[r+1][c] == 1
		}
	case m.nRows - 1:
		switch c {
		case 0:
			return m.values[r-1][c] == 1 || m.values[r][c+1] == 1
		case m.nColumns - 1:
			return m.values[r][c-1] == 1 || m.values[r-1][c] == 1
		default:
			return m.values[r][c-1] == 1 || m.values[r][c+1] == 1 || m.values[r-1][c] == 1
		}
	}
	switch c {
	case 0:
		switch r {
		case 0:
			return m.values[r+1][0] == 1 || m.values[r][c+1] == 1
		case m.nRows - 1:
			return m.values[r][c] == 1 || m.values[r][c+1] == 1
		default:
			return m.values[r-1][c] == 1 || m.values[r][c+1] == 1 || m.values[r+1][c] == 1
		}
	case m.nRows - 1:
		switch c {
		case 0:
			return m.values[r][c-1] == 1 || m.values[r+1][c] == 1
		case m.nColumns - 1:
			return m.values[r][c-1] == 1 || m.values[r-1][c] == 1
		default:
			return m.values[r-1][c] == 1 || m.values[r][c-1] == 1 || m.values[r+1][c] == 1
		}
	}
	return m.values[r][c-1] == 1 || m.values[r-1][c] == 1 || m.values[r][c+1] == 1 || m.values[r+1][c] == 1
}

func (m matrix) daysToSync(days int) int {
	fmt.Println(days)
	m.Print()
	dirty := false
	v := make([][]int, m.nRows)
	for i := 0; i < m.nRows; i++ {
		v[i] = make([]int, m.nColumns)
		for j := 0; j < m.nColumns; j++ {
			if m.canBySynced(i, j) {
				v[i][j] = 1
				dirty = true
				continue
			}
			v[i][j] = 0
		}
	}
	if dirty {
		m.values = v
		days++
		return m.daysToSync(days)
	}
	return days
}

// Given a matrix of 0s and 1s, where 0 means that a server is not synced and 1 means that it is:
// return the number of days that it would take to sync all the servers not synced.
// A synced server can sync any non synced server if it's adjacent (up, down, left, right).
func main() {

	m := newMatrix(16, 16)

	m.daysToSync(0)

}
