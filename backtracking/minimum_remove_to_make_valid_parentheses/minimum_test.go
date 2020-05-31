package minimum

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	for _, testCase := range []struct {
		input    string
		expected string
	}{
		{input: "lee(t(c)o)de)", expected: "lee(t(c)o)de"},
		{input: "a)b(c)d", expected: "ab(c)d"},
		{input: "))((", expected: ""},
		{input: "(a(b(c)d)", expected: "a(b(c)d)"},
	} {
		if res := minRemoveToMakeValid(testCase.input); testCase.expected != res {
			t.Fatalf("expected %q for input %q, got %q instead", testCase.expected, testCase.input, res)
		}
	}
}
