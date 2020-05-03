package main

import (
	"fmt"
	"strings"
)

// T: O(n)
// S: O(1)
func backspaceCompare(S string, T string) bool {
	for {
		if idx := strings.IndexRune(S, '#'); idx != -1 {
			if idx == 0 {
				S = S[1:]
			} else {
				S = S[:idx-1] + S[idx+1:]
			}
			continue
		}
		break
	}

	for {
		if idx := strings.IndexRune(T, '#'); idx != -1 {
			if idx == 0 {
				T = T[1:]
			} else {
				T = T[:idx-1] + T[idx+1:]
			}
			continue
		}
		break
	}

	return S == T

}

func main() {
	fmt.Println(backspaceCompare("ab#c", "ad#c")) // true
	fmt.Println(backspaceCompare("ab##", "c#d#")) // true
	fmt.Println(backspaceCompare("a##c", "#a#c")) // true
	fmt.Println(backspaceCompare("a#c", "b"))     // false
}
