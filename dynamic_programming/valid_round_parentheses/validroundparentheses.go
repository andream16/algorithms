package main

func main() {}

type stack struct {
    parentheses []int
}

func (s *stack) push(p int) {
    s.parentheses = append([]int{p}, s.parentheses...)
}

func (s *stack) pop() {
    if len(s.parentheses) == 0 {
        return
    }
    s.parentheses = s.parentheses[1:]
}

func (s *stack) empty() bool {
    return len(s.parentheses) == 0
}

func (s *stack) peek() int {
    return s.parentheses[0]
}

func longestValidParentheses(s string) int {
    
    var (
        m = 0
        stk = &stack{
            parentheses : []int{-1},
        }
    )
    
    for i, c := range s {
        if c == '(' {
            stk.push(i)
            continue
        }
        stk.pop()
        if stk.empty() {
            stk.push(i)
            continue
        }
        m = max(m, i - stk.peek())
    }
    
    return m

}


func max(n1, n2 int) int {
    if n1 >= n2 {
        return n1
    }
    return n2
}
