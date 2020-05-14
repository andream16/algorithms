package kdigits

import "strings"

// We remove the number that has bigger magnitude on the left of the current one as it will lead to a smaller number:
// arr = [1,2,3,4,5,2,6,4], k = 4. We keep piling up until we find a smaller number on the right. In that case, we
// remove from the stack until the condition is valid and k > 0. A stack facilitates this. If we have other k digits left
// we simply remove them from the end of the result. We also need to remove trailing zeros from results like "002" -> "2".
//
// T: O(n+k) + O(n) + O(n) -> O(n)
// Time Complexity is O(n) because, even if we have a nested loop, we know that we are going to loop a max of k times
// globally and since 0 < k <= n the complexity of this is bound to 2N.
//
// S: O(n)
func removeKdigits(num string, k int) string {
	if k >= len(num) {
		return "0"
	}

	s := &stack{}

	for _, c := range num {
		for k > 0 && !s.empty() && int(s.peak()-'0') > int(c-'0') {
			s.pop()
			k--
		}
		s.push(c)
	}

	// remove extra digits if k>0.
	str := s.toString()[:len(s.nums)-k]

	// removing trailing zeros
	return removeZeros(str)

}

type stack struct {
	nums []rune
}

func (s *stack) push(n rune) {
	s.nums = append(s.nums, n)
}

func (s *stack) pop() {
	s.nums = s.nums[:len(s.nums)-1]
}

func (s *stack) peak() rune {
	return s.nums[len(s.nums)-1]
}

func (s *stack) empty() bool {
	return len(s.nums) == 0
}

func (s *stack) toString() string {
	var builder strings.Builder
	for _, n := range s.nums {
		builder.WriteRune(n)
	}
	return builder.String()
}

func removeZeros(s string) string {
	if len(s) > 1 && s[0] == '0' {
		return removeZeros(s[1:])
	}
	return s
}
