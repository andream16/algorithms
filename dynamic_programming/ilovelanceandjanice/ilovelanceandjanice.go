package main

import (
	"fmt"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
	// a is 97, z is 122
	maxRune = 122
)

func main() {
	// "wrw blf hvv ozhg mrtsg'h vkrhlwv?" -> "did you see last night's episode?"
	// "Yvzs! I xzm'g yvorvev Lzmxv olhg srh qly zg gsv xlolmb!!" -> "Yeah! I can't believe Lance lost his job at the colony!!"
	fmt.Println(solution("Yvzs! I xzm'g yvorvev Lzmxv olhg srh qly zg gsv xlolmb!!"))
}

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
