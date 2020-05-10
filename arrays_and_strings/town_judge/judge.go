package main

import "fmt"

// The problem can be approached like a graph problem. An edge E is identified by a trust pair.
// We have a judge when the number of indegree edges = N - 1. So, this vertex will have all nodes
// going into it except one. We decrease the count when we have an outgoing edge and increase when we have an
// incoming one.
//
// T: O(e)
// S: O(n)
func findJudge(N int, trust [][]int) int {
	if N <= 1 {
		return N
	}

	citizenScores := make(map[int]int, N)

	for _, t := range trust {
		citizenScores[t[0]]--
		citizenScores[t[1]]++
	}

	for citizen, score := range citizenScores {
		if score == N - 1 {
			return citizen
		}
	}

	return -1
}

func main() {
	fmt.Println(findJudge(2, [][]int{{1,2}})) // 2
	fmt.Println(findJudge(3, [][]int{{1,3},{2,3}})) // 3
	fmt.Println(findJudge(3, [][]int{{1,3},{2,3}, {3,1}})) // -1
	fmt.Println(findJudge(3, [][]int{{1,2},{2,3}})) // -1
	fmt.Println(findJudge(4, [][]int{{1,3},{1,4},{2, 3}, {2,4}, {4,3}})) // 3
	fmt.Println(findJudge(4, [][]int{
		{1,3},{3,2},{1, 3}, {4,1}, {3, 1}, {2,1}, {2,3}, {4,3}, {4,2}, {3,4}, {2,4},
	})) // -1
	fmt.Println(findJudge(1, [][]int{})) // 1
}
