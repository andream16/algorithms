package main

import (
	"fmt"
	"sort"
)

type Value struct {
	tile  string
	count int
	value int
	valid bool
}

type ByCount []Value

func (b ByCount) Len() int           { return len(b) }
func (b ByCount) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByCount) Less(i, j int) bool { return b[i].count > b[j].count }

// TODO check again
func minDominoRotations(A []int, B []int) int {

	const (
		tileA = "A"
		tileB = "B"
	)

	var (
		sortedValues     []Value
		aValues, bValues = map[int]int{}, map[int]int{}
		validValues      = map[int]bool{}
	)

	for i, a := range A {
		aCount, _ := aValues[a]
		aCount++
		aValues[a] = aCount
		sortedValues = append(sortedValues, Value{
			tile:  tileA,
			count: aCount,
			value: a,
			valid: validValues[i],
		})
	}

	for i, b := range B {
		bCount, _ := bValues[b]
		bCount++
		bValues[b] = bCount
		sortedValues = append(sortedValues, Value{
			tile:  tileB,
			count: bCount,
			value: b,
			valid: B[i] != A[i],
		})
	}

	sort.Sort(ByCount(sortedValues))

	for _, v := range sortedValues {
		switch v.tile {
		case tileA:
			vb, ok := bValues[v.value]
			if ok {
				if r := len(A) - v.count; vb >= r {
					return r
				}
			}
		case tileB:
			va, ok := aValues[v.value]
			if ok {
				if r := len(B) - v.count; va >= r {
					return r
				}
			}
		}
	}

	return -1

}

func main() {
	fmt.Println(minDominoRotations([]int{2, 1, 2, 4, 2, 2}, []int{5, 2, 6, 2, 3, 2})) // 2
	fmt.Println(minDominoRotations([]int{3, 5, 1, 2, 3}, []int{3, 6, 3, 3, 4}))       // -1
}
