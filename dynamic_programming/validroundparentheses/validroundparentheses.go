type stack struct {
    p []string
}

func (s *stack) push(str string) {
    s.p = append([]string{str}, s.p...)
}

func (s *stack) pop() {
    if s.empty() {
        return
    }
    s.p = s.p[1:]
}

func (s *stack) empty() bool {
    return len(s.p) == 0
}

func (s *stack) reset() {
    s.p = []string{}
}

func longestValidParentheses(s string) int {
    if s == "" {
        return 0
    }
    
    var (
        counter = 0
        longest = 0
        stk = &stack{
            p: []string{},
        }
    )
    
    for _, c := range s {
        fmt.Println(string(c), stk.p)
        cs := string(c)
        if cs == "(" {
            stk.push(cs)
            continue
        }
        stk.pop()
        counter += 2
        if stk.empty() {
            if counter > longest {
                longest = counter
            }
            counter = 0
        }
    }
    
    if counter > longest {
        longest = counter
    }
    
    return longest
    
}
