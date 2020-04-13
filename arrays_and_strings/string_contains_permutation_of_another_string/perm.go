package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	var (
		asBytes, curr int32
		l1            = len(s1)
		nChars        int
		m             = map[int32]int{}
	)

	for _, c := range s1 {
		asBytes += c
		m[c] += 1
	}

	for i := 0; i < len(s2); i++ {
		nChars = 1
		curr = int32(s2[i])
		if curr == asBytes && l1 == nChars {
			return true
		}
		clone := m
		for j := i + 1; j < len(s2); j++ {
			v, ok := clone[int32(s2[j])]
			if !ok || v == 0 {
				break
			}
			clone[int32(s2[j])] -= 1
			nChars++
			if nChars > l1 {
				nChars = 1
				curr = int32(s2[j])
				j = j - 2
				clone = m
				continue
			}
			curr += int32(s2[j])
			if curr == asBytes && nChars == l1 {
				return true
			}
		}
	}
	return false
}

// Input: s1 = "ab" s2 = "eidbaooo"
// Output: True
// Explanation: s2 contains one permutation of s1 ("ba").

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo")) // true
	//fmt.Println(checkInclusion("ab", "eidboaoo"))      // false
	//fmt.Println(checkInclusion("abc", "ccccbbbbaaaa")) // false
}
