package hash

import "testing"

func TestMyHashMap(t *testing.T) {
	hm := Constructor()
	hm.Put(1, 1)
	hm.Put(2, 2)
	if v := hm.Get(1); 1 != v {
		t.Fatalf("expected 1, got %d", v)
	}
	if v := hm.Get(3); -1 != v {
		t.Fatalf("expected -1, got %d", v)
	}
	hm.Put(2, 1)
	if v := hm.Get(2); 1 != v {
		t.Fatalf("expected 1, got %d", v)
	}
	hm.Remove(2)
	if v := hm.Get(2); -1 != v {
		t.Fatalf("expected -1, got %d", v)
	}
}
