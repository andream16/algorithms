package parentheses

// This problem can be solved using backtracking. We know that a string is valid if the number of ( is equal to the
// number of ). The solution of this problem can be better understood using with a whiteboard. View the picture at
// github.com/andream16/assets/images/22_generate_parentheses.jpg.
// Keep in mind `numOpen < max` for right branch and `numClose < numOpen` for left branch.
//
// T: O(b^d) where b is the branching factor and d is the depth.
// 	  In this case it will be O(2^(2*n)) because the branching factor is 2 and the depth is 2*n max as strings are
//    max 2*n long.
// S: Stack is O(b^d)
func generateParenthesis(n int) []string {
	var combs []string
	backtrack(&combs, "", 0, 0, n)
	return combs
}

func backtrack(combs *[]string, curr string, numOpen, numClose, max int) {
	if max*2 == len(curr) {
		*combs = append(*combs, curr)
		return
	}
	if numOpen < max {
		backtrack(combs, curr+"(", numOpen+1, numClose, max)
	}
	if numClose < numOpen {
		backtrack(combs, curr+")", numOpen, numClose+1, max)
	}
}
