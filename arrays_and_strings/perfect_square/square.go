package main

import "fmt"

// T: O(log n) because the sequence converges quadratically.
// S: O(1)
//
// Using Newton's algorithm allows us to solve this in O(log n).
// We start with a reasonable guess, half of the target (e.g. if we have 4, starting with 2 makes sense).
// Then until guess*guess < target, to get the next guess, we calculate the tangent with f(x) and f(x+1) is
// the intersection with the x axis. We keep doing this until guess*guess is bigger than the value or equal.
//
// https://en.wikipedia.org/wiki/Newton%27s_method
func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}
	guess := num / 2
	for guess*guess > num {
		guess = (guess + num/guess) / 2
	}
	return guess*guess == num
}

func main() {
	fmt.Println(isPerfectSquare(16)) // true
	fmt.Println(isPerfectSquare(14)) // false
	fmt.Println(isPerfectSquare(81)) // true
}
