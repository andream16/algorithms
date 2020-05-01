package main

import "fmt"

// TODO study this better
func findStrobogrammatic(n int) []string {
	res := []string{}
	findStroboGrammaticRec(make([]rune, n), 0, n-1, &res)
	return res
}

func findStroboGrammaticRec(s []rune, left, right int, res *[]string) {
	if left > right {
		*res = append(*res, string(s))
		return
	}
	if left == right {
		s[left] = '0'
		*res = append(*res, string(s))
		s[left] = '1'
		*res = append(*res, string(s))
		s[left] = '8'
		*res = append(*res, string(s))
		return
	}
	if left != 0 {
		s[left] = '0'
		s[right] = '0'
		findStroboGrammaticRec(s, left+1, right-1, res)
	}

	s[left] = '1'
	s[right] = '1'
	findStroboGrammaticRec(s, left+1, right-1, res)

	s[left] = '8'
	s[right] = '8'
	findStroboGrammaticRec(s, left+1, right-1, res)

	s[left] = '6'
	s[right] = '9'
	findStroboGrammaticRec(s, left+1, right-1, res)

	s[left] = '9'
	s[right] = '6'
	findStroboGrammaticRec(s, left+1, right-1, res)

}

func main() {
	//fmt.Println(findStrobogrammatic(2))
	fmt.Println(findStrobogrammatic(4))
}
