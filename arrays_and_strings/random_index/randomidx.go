package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	valToRandIdx map[int][]int
}

func Constructor(nums []int) Solution {
	m := map[int][]int{}
	for i, n := range nums {
		if len(m[n]) == 0 {
			m[n] = []int{i}
			continue
		}
		m[n] = append(m[n], i)
	}
	return Solution{
		valToRandIdx: m,
	}
}

func (this *Solution) Pick(target int) int {
	return this.valToRandIdx[target][rand.Intn(len(this.valToRandIdx[target]))]
}

func main() {
	solution := Constructor([]int{1, 2, 3, 3, 3})
	fmt.Println(solution.Pick(1)) // 0
	fmt.Println(solution.Pick(2)) // 1
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
	fmt.Println(solution.Pick(3)) // 2, 3, 4
}
