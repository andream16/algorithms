package wordladder

import (
	"testing"
)

func TestLadderLength(t *testing.T) {
	for _, tc := range []struct{
		beginWord string
		endWord string
		words []string
		expected int
	}{
		{
			beginWord: "hit",
			endWord: "cog",
			words: []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected: 5,
		},
		{
			beginWord: "hit",
			endWord: "cog",
			words: []string{"hot", "dot", "dog", "lot", "log"},
			expected: 0,
		},
	}{
		if length := ladderLength(tc.beginWord, tc.endWord, tc.words); tc.expected != length {
			t.Fatalf("expected %d length, got %d instead", tc.expected, length)
		}
	}
}
