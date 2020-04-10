package main

import (
	"fmt"
	"strconv"
)

// Input: 123
// Output: "One Hundred Twenty Three"

// Input: 12345
// Output: "Twelve Thousand Three Hundred Forty Five"

func numberToWords(num int) string {

	var (
		res    = ""
		digits = 1
		delIdx = 0
		memo   = ""
		s      = strconv.Itoa(num)
	)

	for i := len(s) - 1; i >= 0; i-- {
		cn, _ := strconv.Atoi(string(s[i]))
		switch digits {
		case 1:
			res += getSingle(cn)
		case 10:
			if cn == 1 {
				twoDigits, _ := strconv.Atoi(s[i:])
				res = getTwoDigits(twoDigits)
			} else {
				res = getSecond(cn) + " " + res
			}
		case 100:
			res = getSingle(cn) + " Hundred " + res
		case 1000:
			memo = getSingle(cn) + " Thousand "
			res = memo + res
			delIdx = len(memo)
		case 10000:
			if cn == 1 {
				twoDigits, _ := strconv.Atoi(s[:i+2])
				res = res[delIdx:]
				res = getTwoDigits(twoDigits) + " Thousand " + res
			} else {
				res = getSecond(cn) + " " + res
			}
		}
		digits *= 10
	}
	return res
}

func getSingle(n int) string {
	m := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
		4: "Four",
		5: "Five",
		6: "Six",
		7: "Seven",
		8: "Eight",
		9: "Nine",
	}
	num, ok := m[n]
	if ok {
		return num
	}
	return "Zero"
}

func getSecond(n int) string {
	m := map[int]string{
		2: "Twenty",
		3: "Thirty",
		4: "Forty",
		5: "Fifty",
		6: "Sixty",
		7: "Seventy",
		8: "Eighty",
		9: "Ninety",
	}
	num, ok := m[n]
	if ok {
		return num
	}
	return "Ten"
}

func getTwoDigits(n int) string {
	m := map[int]string{
		11: "Eleven",
		12: "Twelve",
		13: "Thirteen",
		14: "Fourteen",
		15: "Fifteen",
		16: "Sixteen",
		17: "Seventeen",
		18: "Eighteen",
		19: "Nineteen",
	}
	num, ok := m[n]
	if ok {
		return num
	}
	return "Ten"
}

func main() {
	//fmt.Println(numberToWords(123))
	//fmt.Println(numberToWords(18))
	//fmt.Println(numberToWords(999))
	//fmt.Println(numberToWords(10))
	fmt.Println(numberToWords(12345))
}
