package anagrams

import "strings"

// This problem can be solved in many different ways. The most efficient way of solving this is converting each string
// to a format that would look the same for all anagrams of that string, allowing for constant time comparison, instead
// of looping every time. To do so, we convert each string to a string represented by the count of its characters:
// 	- "abbccc" -> "12300000000000000000000000"
// To do so, we need to create a new array of size 26, one for each character. This allows to change the count in
// constant time. With the help of a map, we can understand if an anagram has been visited already or not. Based on that
// we add it to one group or another.
//
// Another implementation I went for was find all anagrams for each passed word. This was extremely not efficient
// unfortunately.
//
// T: O(n*s) where s is the longest string.
// S: O(n*s)
func groupAnagrams(strs []string) [][]string {

	var (
		groups [][]string
		anagramToGroupIndex = map[string]int{}
	)

	for _, str := range strs {
		var (
			countChars = make([]rune, 26)
			b strings.Builder
		)
		for _, char := range str {
			countChars[char - 'a']++
		}
		for i := range countChars {
			b.WriteRune(countChars[i])
		}
		currStr := b.String()
		idx, ok := anagramToGroupIndex[currStr]
		if ok && len(groups) > 0 {
			groups[idx] = append(groups[idx], str)
			continue
		}
		anagramToGroupIndex[currStr] = len(groups)
		groups = append(groups, []string{str})
	}

	return groups

}
