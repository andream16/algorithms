package longestpalindrome

// For each character, we expand from its center to its left & right.
// We keep expanding until we reach the beginning/end of the string or s[left] != s[right].
// Every time we finish analysing one character, we check if the given palindrome is the longest and in case we
// update the answer.
// We check for odd and even length palindromes.
//
// T: O(n^2)
// O: O(1)
func longestPalindrome(s string) string {
	longest := ""

	// Expand at the center and find palindrome string
	// Try to do that to every letter in the string
	for i := 0; i < len(s); i++ {

		// Palindrome string's length could be odd number or even number
		// Example:
		// "aba" -> odd number
		// "abba" -> even number
		odd, even := expandAndCount(s, i, i), expandAndCount(s, i, i+1)

		// Get the longest string by comparing
		longer := ""
		if len(odd) > len(even) {
			longer = odd
		} else {
			longer = even
		}

		if len(longer) > len(longest) {
			longest = longer
		}
	}

	return longest
}

// This function will find the longest string that centered with given letter
// All the indexes, positions, all status will he handled within this function, pretty clean
// Unlike many other solutions with all sorts of status flying around
func expandAndCount(s string, start, end int) string {
	for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
		start -= 1
		end += 1
	}
	// After getting the longest string, start will be added 1 and end will be minus 1 anyway,
	// for slice, start need to be minus 1 and end is exactly what we need
	// so we get this:
	return s[start+1 : end]
}
