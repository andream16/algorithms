package main

import (
	"fmt"
	"strings"
)

func numJewelsInStones(J string, S string) int {
	res := 0
	for _, c := range S {
		if -1 != strings.IndexRune(J, c) {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb")) // 3
	fmt.Println(numJewelsInStones("z", "ZZ"))       // 0
}
