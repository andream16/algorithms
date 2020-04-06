package main

import "fmt"

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	// a is 97, z is 122
	maxRune = 122
)

func solution(s string) string {
	message := make([]rune, len(s))
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			message = append(message, rune(alphabet[maxRune-c]))
			continue
		}
		message = append(message, c)
	}
	return string(message)
}
