package sortlogs

import (
	"bytes"
	"sort"
	"unicode"
)

// While comparing strings, we can keep in mind few rules:
// - if both s1 & s2 don't contain digits, we compare them
// - if s1 contains digits and s2 doesn't, s2 comes first
// - if s1 doesn't contain digits and s2 does, s1 comes first
// - if both s1 & s2 are digits, we don't need to do anything
// Using sort.SliceStable makes sure to order only what has to be without changing original order.
//
// T: O(n log n)
// S: O(1)
func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {

		idx1, idx2 := bytes.IndexRune([]byte(logs[i]), ' '), bytes.IndexRune([]byte(logs[j]), ' ')

		isLetter1, isLetter2 := unicode.IsLetter(rune(logs[i][idx1+1])), unicode.IsLetter(rune(logs[j][idx2+1]))

		switch {
		case isLetter1 && isLetter2:
			return logs[i][idx1+1:] <= logs[j][idx2+1:]
		case isLetter1 && !isLetter2:
			return true
		case !isLetter1 && isLetter2:
			return false
		}
		return false
	})

	return logs
}
