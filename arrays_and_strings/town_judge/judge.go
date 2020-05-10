package main

import "fmt"

// T: O(n) + O(n)
// S: O(n) + O(n*n-1)
func findJudge(N int, trust [][]int) int {
	if N <= 1 {
		return N
	}
	citizens, judges := make(map[int]struct{}, N), make(map[int]map[int]struct{}, N)
	for _, t := range trust {
		_, ok := judges[t[1]]
		if !ok {
			judges[t[1]] = map[int]struct{}{}
		}
		judges[t[1]][t[0]] = struct{}{}
		citizens[t[0]] = struct{}{}
	}

	for j, c := range judges {
		if _, ok := citizens[j]; !ok && len(c) == N-1 {
			return j
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
