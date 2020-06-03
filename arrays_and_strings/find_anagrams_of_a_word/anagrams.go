package anagrams

// We can view this problem like a recursion problem.
// To find the anagrams of a string, we need to iterate on this string's character and form anagrams of length n from
// what remains in the string, excluded the current character.
//
// Example for 'eat':
// 	- 1) we begin with the empty string '' and the remaining string is 'eat'
// 	- 2) we consider 'e' with 'at' remaining, 'a' with 'et' remaining and 't' with 'ea' remaining
//  - 3) we keep going like this until the string has the correct length:
//       - [('ea', 't'), ('et', 'a'), ('ae', 't'), ('at', 'e'), ('te', 'a'), ('ta', 'e')]
//       - ['eat', 'eta', 'aet', 'ate', 'tea', 'tae']
//
// T: O(n*s) where s is the length of the remaining chars
// S: O(n*s)
func findAnagrams(w string) []string {
	return perm("", w, len(w), []string{})
}

func perm(str, rem string, n int, res []string) []string {
	if n == len(str) {
		return append(res, str)
	}
	for i := 0; i < len(rem); i++ {
		res = perm(str+string(rem[i]), rem[:i]+rem[i+1:], n, res)
	}
	return res
}
