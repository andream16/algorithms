package lru_cache_generic

import "testing"

func TestLRUCache_Get(t *testing.T) {
	t.Run("ints", func(t *testing.T) {
		cache := New[int, int](2)
		cache.Put(1, 1)
		cache.Put(2, 2)

		res, ok := cache.Get(1)
		expects(t, true, 1, ok, res)

		cache.Put(3, 3)
		res, ok = cache.Get(2)
		expects(t, false, 0, ok, res)

		cache.Put(4, 4)
		res, ok = cache.Get(1)
		expects(t, false, 0, ok, res)

		res, ok = cache.Get(3)
		expects(t, true, 3, ok, res)

		res, ok = cache.Get(4)
		expects(t, true, 4, ok, res)
	})
	t.Run("strings", func(t *testing.T) {
		cache := New[string, int](2)
		cache.Put("1", 1)
		cache.Put("2", 2)

		res, ok := cache.Get("1")
		expects(t, true, 1, ok, res)

		cache.Put("3", 3)
		res, ok = cache.Get("2")
		expects(t, false, 0, ok, res)

		cache.Put("4", 4)
		res, ok = cache.Get("1")
		expects(t, false, 0, ok, res)

		res, ok = cache.Get("3")
		expects(t, true, 3, ok, res)

		res, ok = cache.Get("4")
		expects(t, true, 4, ok, res)
	})
}

func expects(
	t *testing.T,
	expectedFound bool,
	expectedVal any,
	actualFound bool,
	actualVal any,
) {
	t.Helper()
	if expectedFound != actualFound {
		t.Fatalf("expected to be found: %v, got: %v", expectedFound, actualFound)
	}
	if expectedVal != actualVal {
		t.Fatalf("expected %v, got %v", expectedVal, actualVal)
	}
}
