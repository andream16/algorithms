package main

import "fmt"

// T: O(D*L*C), D = |digits|, L = |letters| per Digit, C = |previous combinations|
// O(23) -> O(3^3, 3^3)
// S: O(L^N)
func letterCombinations(digits string) []string {
	var (
		combs   = []string{}
		letters = []string{}
	)
	for _, d := range digits {
		letters = getLetters(d)
		if len(combs) == 0 {
			combs = append(combs, letters...)
			continue
		}
		newComb := []string{}
		for _, c := range combs {
			for _, l := range letters {
				newComb = append(newComb, c+l)
			}
		}
		combs = newComb
	}
	return combs
}

func getLetters(b rune) []string {
	if letters, ok := map[rune][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}[b]; ok {
		return letters
	}
	return nil
}

// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

func main() {
	fmt.Println(letterCombinations("23"))
}
