package minimum

import "strings"

type stack struct {
	indexes []int
}

func (s *stack) push(i int) {
	s.indexes = append([]int{i}, s.indexes...)
}

func (s *stack) pop() int {
	i := s.indexes[0]
	s.indexes = s.indexes[1:]
	return i
}

func (s *stack) empty() bool {
	return 0 == len(s.indexes)
}

// To solve this problem, we rely on a stack. We add do the stack the indexes of the respective characters that we're
// considering.
//
// We add to the stack every time we find a '('. When we find a ')', we mark it for deletion if the stack is empty
// (since it means that the number of closing parentheses does not match the number of opening parentheses) otherwise,
// we pop the top of the stack to say that a ')' matches a '('.
//
// After this greedy phase has concluded, we check if the stack is empty. If so, we are done, otherwise, we need to mark
// for deletion all the leftover '('.
//
// Now we write the final result by skipping the characters at the current index reported for deletion.
//
// T: O(n)
// S: O(n)
func minRemoveToMakeValid(s string) string {
	var (
		stck        = &stack{}
		toBeRemoved = map[int]struct{}{}
		b           strings.Builder
	)

	for i, r := range s {
		switch r {
		case ')':
			if stck.empty() {
				toBeRemoved[i] = struct{}{}
				continue
			}
			stck.pop()
		case '(':
			stck.push(i)
		}
	}

	for _, i := range stck.indexes {
		toBeRemoved[i] = struct{}{}
	}

	for i, r := range s {
		if _, ok := toBeRemoved[i]; ok {
			continue
		}
		b.WriteRune(r)
	}

	return b.String()
}
