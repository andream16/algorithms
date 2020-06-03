package validanagram

import "strings"

// To have a constant comparison between the strings, we convert them to the same format consisting of a string
// formed by adding how many times we found a certain character.
//
// Example:
// - "abbccc" -> "12300000000000000000000000", "cbacbc" -> "12300000000000000000000000"
//
// T: O(n)
// S: O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return getAnagramString(s) == getAnagramString(t)
}

func getAnagramString(s string) string {
	var (
		b strings.Builder
		dictCount = make([]rune, 26)
	)

	for _, r := range s {
		dictCount[r-'a']++
	}

	for _, r := range dictCount {
		b.WriteRune(r)
	}

	return b.String()
}
