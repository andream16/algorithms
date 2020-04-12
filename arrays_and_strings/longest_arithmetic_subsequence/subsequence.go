package main

import "fmt"

func longestArithSeqLength(A []int) int {
	if l := len(A); l <= 2 {
		return l
	}

	var (
		longest, currSub []int
		curr, next, diff int
	)

	for i := 0; i < len(A)-1; i++ {
		curr = A[i]
		for k := i + 1; k < len(A); k++ {
			next = A[k]
			diff = curr - next
			currSub = []int{curr, next}
			for j := k + 1; j < len(A); j++ {
				if diff == next-A[j] {
					currSub = append(currSub, A[j])
					next = A[j]
				}
			}
			if len(currSub) > len(longest) {
				longest = currSub
			}
		}
	}

	return len(longest)

}

//Input: [20,1,15,3,10,5,8]
//Output: 4
//Explanation:
//The longest arithmetic subsequence is [20,15,10,5].
func main() {
	fmt.Println(longestArithSeqLength([]int{3, 6, 9, 12}))            // 4
	fmt.Println(longestArithSeqLength([]int{20, 1, 15, 3, 10, 5, 8})) // 4
	fmt.Println(longestArithSeqLength([]int{9, 4, 7, 2, 10}))         // 3
}
