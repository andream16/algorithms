package frequent

import "testing"

func TestTopKFrequent(t *testing.T) {
	for _, testCase := range []struct {
		words    []string
		k        int
		expected []string
	}{
		{
			words:    []string{"i", "love", "leetcode", "i", "love", "coding"},
			k:        2,
			expected: []string{"i", "love"},
		},
		{
			words:    []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"},
			k:        4,
			expected: []string{"the", "is", "sunny", "day"},
		},
	} {
		res := topKFrequent(testCase.words, testCase.k)
		if testCase.k != len(res) {
			t.Fatalf("expected %d results, got %d instead", testCase.k, len(res))
		}
		for i := 0; i < testCase.k; i++ {
			if testCase.expected[i] != res[i] {
				t.Fatalf("expected %q at position %d, got %q instead", testCase.expected[i], i, res[i])
			}
		}
	}
}
