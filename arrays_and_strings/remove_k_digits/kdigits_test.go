package kdigits

import "testing"

func TestRemoveKDigits(t *testing.T) {
	t.Run("it should return 1219", func(t *testing.T) {
		if num := removeKdigits("1432219", 3); num != "1219" {
			t.Fatalf("expected 1219, got %s", num)
		}
	})
	t.Run("it should return 200", func(t *testing.T) {
		if num := removeKdigits("10200", 1); num != "200" {
			t.Fatalf("expected 200, got %s", num)
		}
	})
	t.Run("it should return 0", func(t *testing.T) {
		if num := removeKdigits("10", 2); num != "0" {
			t.Fatalf("expected 0, got %s", num)
		}
	})
	t.Run("it should return 0", func(t *testing.T) {
		if num := removeKdigits("10", 1); num != "0" {
			t.Fatalf("expected 0, got %s", num)
		}
	})
	t.Run("it should return 11", func(t *testing.T) {
		if num := removeKdigits("112", 1); num != "11" {
			t.Fatalf("expected 11, got %s", num)
		}
	})
}
