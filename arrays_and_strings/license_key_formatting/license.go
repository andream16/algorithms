package main

import (
	"fmt"
)

func licenseKeyFormatting(S string, K int) string {
	if K == 0 {
		return ""
	}
	var (
		s       string
		counter int
	)
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}
		if counter == K {
			s = "-" + s
			counter = 0
		}
		counter++
		if S[i] >= 'a' && S[i] <= 'z' {
			s = string(S[i]-32) + s
			continue
		}
		s = string(S[i]) + s
	}
	return s
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
