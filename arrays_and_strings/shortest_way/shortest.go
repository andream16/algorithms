package main

import (
	"fmt"
)

// TODO UNDERSTAND THIS
func shortestWay(source string, target string) int {
	count := 0

	for it := 0; it < len(target); {
		found := false
		for is := 0; is < len(source) && it < len(target); is++ {
			if source[is] == target[it] {
				found = true
				it += 1
			}
		}
		if !found {
			return -1
		}
		count += 1
	}

	return count
}

func main() {
	fmt.Println(shortestWay("xyz", "xzyxz"))              // 3
	fmt.Println(shortestWay("abc", "abcbc"))              // 2
	fmt.Println(shortestWay("abc", "acdbc"))              // -1
	fmt.Println(shortestWay("adbsc", "addddddddddddsbc")) // 13
}
