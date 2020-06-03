package anagrams

import "testing"

func TestFindAnagrams(t *testing.T) {
	expected := []string{"eat", "eta", "aet", "ate",  "tea",  "tae"}
	anagrams := findAnagrams("eat")
	if 6 != len(anagrams) {
		t.Fatalf("expected 6 anagrams for 'eat', found %d", len(anagrams))
	}
	for i:=0; i<len(anagrams); i++ {
		if expected[i] != anagrams[i] {
			t.Fatalf("expected %q at position %d, got %q", expected[i], i, anagrams[i])
		}
	}
}
