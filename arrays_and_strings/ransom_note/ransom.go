package main

import (
	"fmt"
)

// T: O(n)
// S: O(m)
func canConstruct(ransomNote string, magazine string) bool {
	return canConstructIterative(ransomNote, magazine)
}

func canConstructIterative(note, magazine string) bool {
	noteTokens := map[rune]int{}
	for _, n := range note {
		noteTokens[n]++
	}

	for _, m := range magazine {
		n, ok := noteTokens[m]
		if ok {
			if n == 1 {
				delete(noteTokens, m)
				continue
			}
			noteTokens[m]--
		}
	}

	return len(noteTokens) == 0

}
func main() {
	fmt.Println(canConstruct("a", "b"))    // false
	fmt.Println(canConstruct("aa", "ab"))  // false
	fmt.Println(canConstruct("aa", "aab")) // true
}
