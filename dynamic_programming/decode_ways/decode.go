package main

// O(2^n) with no memoization
// O(n) with memoization

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	return numDecodingsRec(s, map[string]int{})
}

func numDecodingsRec(s string, m map[string]int) int {
	if s == "" {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	r, ok := m[s]
	if ok {
		return r
	}
	res := numDecodingsRec(s[1:], m)
	if len(s) > 1 {
		if n, _ := strconv.Atoi(s[0:2]); n <= 26 {
			res += numDecodingsRec(s[2:], m)
		}
	}
	m[s] = res
	return res
}

func main() {
	fmt.Println(numDecodings("2264"))
}
