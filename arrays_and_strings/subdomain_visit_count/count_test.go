package count

import "testing"

func TestSubdomainVisits(t *testing.T) {
	for _, testCase := range []struct {
		input    []string
		expected []string
	}{
		{
			input:    []string{"9001 discuss.leetcode.com"},
			expected: []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"},
		},
		{
			input:    []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
			expected: []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"},
		},
	} {
		res := subdomainVisits(testCase.input)
		if len(testCase.expected) != len(res) {
			t.Fatalf("expected result of length %d, got length %d", len(testCase.expected), len(res))
		}
		for _, r := range res {
			if !contains(r, testCase.expected) {
				t.Fatalf("expected %s result to be found in %v", r, testCase.expected)
			}
		}
	}
}

func contains(needle string, haystack []string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}
	return false
}
