package validpalindromeii

// A string is a valid palindrome when it looks the same reversed. In this case, we can remove one character at a time
// to get a palindrome.
// In the case the string has 0 or 1 character, it's always a palindrome.
// In the other cases, we use two pointers right and left. Right starts at the beginning of the string and left starts
// at the end of the same string. We keep incrementing and decrementing the pointers until they meet. When they meet, we
// know that the left side of the string might look like the right side of it. If that's true, we return true too.
// In the other case, we might have odd length strings and need to check if the left side or the right side are palindrome.
//
// T: O(N)
// S: O(1)
func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1

	for left <= right && s[left] == s[right] {
		left++
		right--
	}

	if left >= right {
		return true
	}

	return isPalindrome(s[:left] + s[left+1:]) || isPalindrome(s[:right] + s[right+1:])

}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}
	return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
}
