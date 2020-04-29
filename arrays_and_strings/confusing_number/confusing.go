package main

import (
	"fmt"
	"strconv"
)

// T: O(n)
// S: O(1)
func confusingNumber(N int) bool {
	s := ""
	for _, n := range strconv.Itoa(N) {
		rs, ok := isConfusing(n)
		if !ok {
			return false
		}
		s = rs + s
	}
	ns, _ := strconv.Atoi(s)
	return N != ns
}

func isConfusing(r rune) (string, bool) {
	v, ok := map[rune]string{
		'0': "0",
		'1': "1",
		'6': "9",
		'8': "8",
		'9': "6",
	}[r]
	return v, ok
}

func main() {
	fmt.Println(confusingNumber(6))  // true
	fmt.Println(confusingNumber(89)) // true
	fmt.Println(confusingNumber(11)) // false
	fmt.Println(confusingNumber(25)) // false
}
