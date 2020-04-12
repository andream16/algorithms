package main

import "fmt"

// Input: s = "eceba", k = 2
// Output: 3
// Explanation: T is "ece" which its length is 3.

// Input: s = "aa", k = 1
// Output: 2
// Explanation: T is "aa" which its length is 2.

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 || s == "" {
		return 0
	}
	return helper(s, "", "", k, 0, map[byte]int{})
}

func helper(s, currS, longest string, k, idx int, currDistinct map[byte]int) int {
	if len(longest) < len(currS) {
		longest = currS
	}
	if idx == len(s) {
		return len(longest)
	}
	c := s[idx]
	if _, ok := currDistinct[c]; ok {
		if idx > 0 && s[idx] != s[idx-1] {
			currDistinct[c] = idx
		}
		currS += string(c)
		return helper(s, currS, longest, k, idx+1, currDistinct)
	}
	if k > len(currDistinct) {
		currS += string(c)
		currDistinct[c] = idx
		return helper(s, currS, longest, k, idx+1, currDistinct)
	}

	oldIdx, _ := currDistinct[currS[0]]
	currS = s[oldIdx:idx] + string(c)

	newCurrDistinct := map[byte]int{}

	for i := range currS {
		newCurrDistinct[currS[i]] = len(s[oldIdx:]) + 1
	}
	return helper(s, currS, longest, k, idx+1, currDistinct)
}

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))       // 3
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))          // 2
	fmt.Println(lengthOfLongestSubstringKDistinct("ab", 1))          // 1
	fmt.Println(lengthOfLongestSubstringKDistinct("bacc", 2))        // 3
	fmt.Println(lengthOfLongestSubstringKDistinct("cdaba", 3))       // 4
	fmt.Println(lengthOfLongestSubstringKDistinct("ababacccccc", 2)) // 7
}
