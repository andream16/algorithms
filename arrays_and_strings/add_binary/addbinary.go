package main

import (
	"fmt"
	"strconv"
)

func main() {}

func addBinary(a string, b string) string {
	var sum string
	helper(a, b, 0, &sum)
	return sum
}

func helper(s1, s2 string, carry int, res *string) {
	if s1 == "" && s2 == "" {
		if carry > 0 {
			*res = strconv.Itoa(carry) + *res
		}
		return
	}

	l1, l2 := len(s1)-1, len(s2)-1
	n1, n2 := 0, 0

	if l1 >= 0 {
		n1, _ = strconv.Atoi(string(s1[l1]))
		s1 = s1[:l1]
	}
	if l2 >= 0 {
		n2, _ = strconv.Atoi(string(s2[l2]))
		s2 = s2[:l2]
	}

	curr := n1 ^ n2 ^ carry
	carry = n1 & n2

	fmt.Println(n1, n2, curr, carry)

	*res = strconv.Itoa(curr) + *res

	helper(s1, s2, carry, res)
}
