package parentheses

import "testing"

func TestGenerateParenthesis(t *testing.T) {
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	res := generateParenthesis(3)
	if l := len(res); 5 != l {
		t.Fatalf("expected 5 results, got %d", len(res))
	}
	for i := range res {
		if expected[i] != res[i] {
			t.Fatalf("expected %s at position %d, found %s instead", expected[i], i, res[i])
		}
	}
}
