package trie

import "testing"

func TestTrie(t *testing.T) {
	tr := Constructor()
	tr.Insert("apple")
	if !tr.Search("apple") {
		t.Fatalf("expected 'apple' to be found in trie")
	}
	if tr.Search("app") {
		t.Fatalf("expected 'app' to not be found in trie")
	}
	if !tr.StartsWith("app") {
		t.Fatalf("expected a word with prefix 'app' to be found in trie")
	}
	tr.Insert("app")
	if !tr.Search("app") {
		t.Fatalf("expected 'app' to be found in trie")
	}
}
