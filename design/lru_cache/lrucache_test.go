package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	t.Run("it should execute the scenario as expected", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		if res := cache.Get(1); 1 != res {
			t.Fatalf("expected 1, got %d for key 1", res)
		}
		cache.Put(3, 3)
		if res := cache.Get(2); -1 != res {
			t.Fatalf("expected -1, got %d for key 2", res)
		}
		cache.Put(4, 4)
		if res := cache.Get(1); -1 != res {
			t.Fatalf("expected -1, got %d for key 1", res)
		}
		if res := cache.Get(3); 3 != res {
			t.Fatalf("expected 3, got %d for key 3", res)
		}
		if res := cache.Get(4); 4 != res {
			t.Fatalf("expected 4, got %d for key 4", res)
		}
	})
}
