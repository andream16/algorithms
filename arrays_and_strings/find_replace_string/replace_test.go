package replacestr

import "testing"

func TestFindReplaceString(t *testing.T) {
	for _, testCase := range []struct {
		givenString  string
		givenIndexes []int
		givenSources []string
		givenTargets []string
		expected     string
	}{
		{
			givenString:  "abcd",
			givenIndexes: []int{0, 2},
			givenSources: []string{"a", "cd"},
			givenTargets: []string{"eee", "ffff"},
			expected:     "eeebffff",
		},
		{
			givenString:  "abcd",
			givenIndexes: []int{0, 2},
			givenSources: []string{"ab", "ec"},
			givenTargets: []string{"eee", "ffff"},
			expected:     "eeecd",
		},
		{
			givenString:  "vmokgggqzp",
			givenIndexes: []int{3, 5, 1},
			givenSources: []string{"kg", "ggq", "mo"},
			givenTargets: []string{"s", "so", "bfr"},
			expected:     "vbfrssozp",
		},
	} {
		res := findReplaceString(testCase.givenString, testCase.givenIndexes, testCase.givenSources, testCase.givenTargets)
		if testCase.expected != res {
			t.Fatalf("expected %q but got %q", testCase.expected, res)
		}
	}
}
