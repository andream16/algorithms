package main

import (
	"fmt"
	"math"
	"strconv"
)

func numDecodings(s string) int {
	return numDecodingsRec(s, map[string]int{}) % int(math.Pow(float64(10), 9)+7)
}

func numDecodingsRec(s string, m map[string]int) int {
	if s == "" {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if s == "**" {
		return 96
	}
	if s[0] == '*' {
		return 9
	}
	r, ok := m[s]
	if ok {
		return r
	}
	res := numDecodingsRec(s[1:], m)
	if len(s) > 1 {
		if n, _ := strconv.Atoi(s[0:2]); n > 0 && n <= 26 {
			res += numDecodingsRec(s[2:], m)
		} else if s[1] == '*' {
			switch s[0] {
			case '1':
				res += 9
			case '2':
				res += 9
			}
		}
	}
	m[s] = res
	return res
}

func main() {
	//fmt.Println(numDecodings("*"))
	//fmt.Println(numDecodings("1*"))
	fmt.Println(numDecodings("4*"))
	//fmt.Println(numDecodings("**"))
}
