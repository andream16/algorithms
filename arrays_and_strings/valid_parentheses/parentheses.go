package parentheses

// We can use a stack to easily solve this problem. If we find an opening bracket, we simply add it to the stack.
// If we find a closing bracket, this has to match with the previous opening one.
// If we find a closing bracket but the stack is empty, it means that the string is not valid.
// Finally, the string is valid only if, at the end, the stack is empty as all parentheses match.
//
// T: O(n)
// S: O(n)
func isValid(s string) bool {
	var stack []rune

	for _, r := range s {
		if isOpen(r) {
			stack = append(stack, r)
			continue
		}
		if 0 == len(stack) || false == matches(stack[len(stack)-1], r) {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func isOpen(r rune) bool {
	return '(' == r || '[' == r || '{' == r
}

func matches(op, cl rune) bool {
	switch cl {
	case ')':
		return op == '('
	case ']':
		return op == '['
	case '}':
		return op == '{'
	}
	return false
}
