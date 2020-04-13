package main

import "fmt"

// Input: words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
// Output: true
// Explanation: As 'h' comes before 'l' in this language, then the sequence is sorted.

func isAlienSorted(words []string, order string) bool {

	mOrder := make([]int, 26)
	for idx, l := range order {
		mOrder[l-'a'] = idx
	}

	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		for k := 0; k < shortestWordLength(word1, word2); k++ {
			l1, l2 := word1[k], word2[k]
			if mOrder[l1-'a'] > mOrder[l2-'a'] {
				return false
			}
		}
		if len(word1) > len(word2) {
			return false
		}
	}

	return true

}

func shortestWordLength(w1, w2 string) int {
	l1, l2 := len(w1), len(w2)
	if l1 < l2 {
		return l1
	}
	return l2
}

func main() {
	fmt.Println(isAlienSorted([]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz")) // true
	fmt.Println(isAlienSorted([]string{"kuvp", "q"}, "ngxlkthsjuoqcpavbfdermiywz"))         // true
}
