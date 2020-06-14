package longest

import "strings"

// To solve this problem, we have to iterate all over the string.
// We have a concept of longest string and current word.
// We keep appending to the current word if the current character is not in it.
// When instead we find an existing character, we update the longest string value and update the current word
// to start from the index after the duplicate.
//
// T: O(n)
// S: O(n)
func lengthOfLongestSubstring(s string) int {

	switch {
	case s == "":
		return 0
	case len(s) == 1:
		return 1
	}

	var (
		currWord string
		longest  int
	)

	for i, c := range s {
		cs := string(c)
		if i == 0 {
			currWord += cs
			continue
		}

		idx := strings.Index(currWord, cs)
		if idx == -1 {
			currWord += cs
			if i != len(s)-1 {
				continue
			}
		}

		if len(currWord) > longest {
			longest = len(currWord)
		}
		currWord = currWord[idx+1:] + cs
	}

	return longest
}
