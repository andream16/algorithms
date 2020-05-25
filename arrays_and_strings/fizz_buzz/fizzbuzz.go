package fizzbuzz

import "fmt"

// T: O(n)
// S: O(n)
func fizzBuzz(n int) []string {

	const (
		fizz     = "Fizz"
		buzz     = "Buzz"
		fizzBuzz = fizz + buzz
	)

	res := make([]string, n)

	for i := 0; i < n; i++ {
		switch {
		case (i+1)%3 == 0 && (i+1)%5 == 0:
			res[i] = fizzBuzz
		case (i+1)%3 == 0:
			res[i] = fizz
		case (i+1)%5 == 0:
			res[i] = buzz
		default:
			res[i] = fmt.Sprintf("%d", i+1)
		}
	}

	return res

}
