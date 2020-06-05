package wordbreak

// This problem can be solved by solving its sub-problems.
// Given 'code' and a dictionary ['c', 'od', 'e', 'y']
// We check if any sub-word in the string exists in the dictionary.
// Based on that, we update a map for future lookups:
// - 'code' -> 'c' & wordBreak('ode'), 'co' & wordBreak('de'), 'cod' & wordBreak('e'), 'code' & wordBreak('')
// - Then, just for the first in this example:
//   - 'ode' -> 'o' & wordBreak('de')
//
// Video with explanation https://www.youtube.com/watch?v=hLQYQ4zj0qg
//
// T: O(n^2) as the recursion tree can grow up to n^2
// S: O(n)
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]struct{}, len(wordDict))

	for _, w := range wordDict {
		dict[w] = struct{}{}
	}

	return wordBreakRec(s, dict, 0, map[int]bool{})
}

func wordBreakRec(s string, dict map[string]struct{}, start int, memo map[int]bool) bool {
	if start == len(s) {
		return true
	}
	if v, ok := memo[start]; ok {
		return v
	}
	for end := start+1; end <= len(s); end++ {
		if _, ok := dict[s[start:end]]; ok && wordBreakRec(s, dict, end, memo) {
			memo[start] = true
			return true
		}
	}
	memo[start] = false
	return false
}
