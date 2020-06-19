package atoi

import (
	"strconv"
	"strings"
)

var (
	nums = map[string]struct{}{
		"0": {},
		"1": {},
		"2": {},
		"3": {},
		"4": {},
		"5": {},
		"6": {},
		"7": {},
		"8": {},
		"9": {},
	}
	signs = map[string]int{
		"+": 1,
		"-": -1,
	}
)

// TODO REVIEW
func myAtoi(str string) int {

	str = strings.Trim(str, " ")
	if str == "" {
		return 0
	}

	firstChar := string(str[0])

	if !isNumberOrSign(firstChar) {
		return 0
	}

	var (
		sign = getSign(firstChar)
		ns   = ""
	)

	if !isNumber(firstChar) {
		str = str[1:]
	}

	for _, c := range str {
		cs := string(c)
		if !isNumber(cs) {
			break
		}
		ns += cs
	}
	if ns == "" {
		return 0
	}

	i, err := strconv.ParseInt(ns, 10, 32)
	if err != nil {
		if sign == -1 {
			return 2147483648 * sign
		}
		return 2147483647
	}
	return int(i) * sign
}

func isNumberOrSign(s string) bool {
	_, ok := nums[s]
	if ok {
		return true
	}
	_, ok = signs[s]
	return ok
}

func isNumber(s string) bool {
	_, ok := nums[s]
	return ok
}

func getSign(s string) int {
	v, ok := signs[s]
	if !ok {
		return 1
	}
	return v
}
