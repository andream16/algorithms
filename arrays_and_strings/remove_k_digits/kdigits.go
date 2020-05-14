package kdigits

// We remove the number that has bigger magnitude on the left of the current one as it will lead to a smaller number:
// arr = [1,2,3,4,5,2,6,4], k = 4. We keep piling up until we find a smaller number on the right. In that case, we
// remove from the stack until the condition is valid and k > 0. A stack facilitates this. If we have other k digits left
// we simply remove them from the end of the result. We also need to remove trailing zeros from results like "002" -> "2".
//
// T: O(n+k) + O(n) -> O(n)
// Time Complexity is O(n) because, even if we have a nested loop, we know that we are going to loop a max of k times
// globally and since 0 < k <= n the complexity of this is bound to 2N.
//
// S: O(n)
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	var str string

	// we are using str like a stack.
	for _, c := range num {
		// keep removing until k > 0 and the criteria matches.
		for k > 0 && str != "" && int(str[len(str)-1]-'0') > int(c-'0') {
			str = str[:len(str)-1]
			k--
		}
		str += string(c)
	}

	// remove extra digits if k>0.
	str = str[:len(str)-k]

	// removing trailing zeros
	return removeZeros(str)

}

func removeZeros(s string) string {
	if len(s) > 1 && s[0] == '0' {
		return removeZeros(s[1:])
	}
	return s
}
