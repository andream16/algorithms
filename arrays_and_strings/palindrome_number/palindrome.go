package palindrome

// As we can't convert to string, we need to find a different way to approach the problem. First thing first, we know
// that negative numbers are never palindromic. In that case we return false. Same thing for numbers that % 10 are
// equal to 0 (10, 100, 1000, ...), they are never palindromic, so we return false in that case too.
// For the other ones, we can simply use an helper reverted number that is going to add up numbers from the other
// way around of the number, from the smallest digit. To do so, we use r = (r * 10) + (x % 10). Now, we can simply
// compare the original input with the reverted one. The only problem is if the original number is formed by
// an odd number of digits. In that case, we divide the reverted by another 10 and compare it. This is fine because the
// digit in the middle is always palindromic.
//
// T: O(log n) - We divided the input by 10 for every iteration, so the time complexity is O(log10(n))
// S: O(1)
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var revNum int

	for revNum < x {
		revNum = revNum*10 + x%10
		x /= 10
	}

	return x == revNum || x == revNum/10

}
