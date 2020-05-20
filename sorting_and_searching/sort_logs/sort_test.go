package sortlogs

import "testing"

func TestReorderLogFiles(t *testing.T) {
	expected := []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}
	res := reorderLogFiles([]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"})

	if len(expected) != len(res) {
		t.Fatalf("expected %d results, got %d", len(expected), len(res))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != res[i] {
			t.Fatalf("expected log at position %d to be %q, got %q", i, expected[i], res[i])
		}
	}
}
