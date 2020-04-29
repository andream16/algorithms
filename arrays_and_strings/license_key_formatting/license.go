package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(S string, K int) string {
	if K == 0 {
		return ""
	}

	var (
		res   = ""
		s     = filterDashes(S)
		start = 0
	)

	if r := len(s) % K; r != 0 {
		start = r
		res += s[:r]
		if len(s) > 1 {
			res += "-"
		}
	}

	for i := start; i < len(s); i += K {
		res += s[i : i+K]
		if i+K < len(s) {
			res += "-"
		}
	}
	return res
}

func filterDashes(s string) string {
	var res string
	for _, c := range s {
		if c != '-' {
			res += strings.ToUpper(string(c))
		}
	}
	return res
}

// Input: S = "5F3Z-2e-9-w", K = 4
//
//Output: "5F3Z-2E9W"

func main() {
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4)) // "5F3Z-2E9W"
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))    // "2-5G-3J"
	fmt.Println(licenseKeyFormatting("2-4A0r7-4k", 3))  // "24-A0R-74K"
	fmt.Println(licenseKeyFormatting("2", 2))           // "2"
}
