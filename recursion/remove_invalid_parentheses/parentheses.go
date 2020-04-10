package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	emptyRes := make([]string, 0)
	if s == "" {
		return emptyRes
	}
	if len(s) == 1 {
		if isChar(s[0]) {
			return []string{s}
		}
	}
	res := removeInvalidParenthesesRec(s, 0, map[string]struct{}{}, []string{})
	if isValid(s) {
		res = append(res, s)
	}
	if len(res) == 0 {
		return emptyRes
	}
	return res
}

func removeInvalidParenthesesRec(s string, idx int, m map[string]struct{}, res []string) []string {
	if idx == len(s)-1 {
		_, ok := m[s[:idx]]
		if !ok && isValid(s[:idx]) {
			res = append(res, s[:idx])
		}
		return res
	}

	if isChar(s[idx]) {
		return removeInvalidParenthesesRec(s, idx+1, m, res)
	}

	tmp := s[:idx] + s[idx+1:]
	_, ok := m[tmp]
	if !ok && isValid(tmp) {
		res = append(res, tmp)
		m[tmp] = struct{}{}
	}

	return removeInvalidParenthesesRec(s, idx+1, m, res)
}

func isValid(s string) bool {
	numOpen, numClose := 0, 0
	for i, c := range s {
		if c == '(' && (i == 0 || i < len(s)-1) {
			numOpen++
		} else if c == ')' {
			numClose++
		}
	}
	return numOpen == numClose
}

func isChar(s byte) bool {
	return s != '(' && s != ')'
}

func main() {
	fmt.Println(removeInvalidParentheses("()())()"))
	fmt.Println(removeInvalidParentheses("(a)())()"))
	fmt.Println(removeInvalidParentheses("n"))
	fmt.Println(removeInvalidParentheses(")("))
	fmt.Println(removeInvalidParentheses("("))
	fmt.Println(removeInvalidParentheses("x("))
	fmt.Println(removeInvalidParentheses("()"))
}
