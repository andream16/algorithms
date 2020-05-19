package parentheses

// We can use a stack to easily solve this problem. If we find an opening bracket, we simply add it to the stack.
// If we find a closing bracket, this has to match with the previous opening one.
// If we find a closing bracket but the stack is empty, it means that the string is not valid.
// Finally, the string is valid only if, at the end, the stack is empty as all parentheses match.
//
// T: O(n)
// S: O(n)
func isValid(s string) bool {
	var stack []byte

	for i := 0; i < len(s); i++ {
		if isOpen(s[i]) {
			stack = append([]byte{s[i]}, stack...)
			continue
		}
		if len(stack) == 0 || !matches(s[i], stack[0]) {
			return false
		}
		stack = stack[1:]
	}

	return len(stack) == 0
}

func isOpen(r byte) bool {
	return r == '(' || r == '[' || r == '{'
}

func matches(closed, open byte) bool {
	switch closed {
	case ')':
		return open == '('
	case ']':
		return open == '['
	case '}':
		return open == '{'
	}
	return false
}
