package add

import "testing"

func TestAddStrings(t *testing.T) {
	t.Run("should sum 33 and 66", func(t *testing.T) {
		if res := addStrings("33", "66"); "99" != res {
			t.Fatalf("expected 99, got %s", res)
		}
	})
	t.Run("should sum 33 and 67", func(t *testing.T) {
		if res := addStrings("33", "67"); "100" != res {
			t.Fatalf("expected 100, got %s", res)
		}
	})
	t.Run("should sum 0 and 67", func(t *testing.T) {
		if res := addStrings("0", "67"); "67" != res {
			t.Fatalf("expected 67, got %s", res)
		}
	})
	t.Run("should sum 67 and 0", func(t *testing.T) {
		if res := addStrings("67", "0"); "67" != res {
			t.Fatalf("expected 67, got %s", res)
		}
	})
	t.Run("should sum 0 and 0", func(t *testing.T) {
		if res := addStrings("0", "0"); "0" != res {
			t.Fatalf("expected 0, got %s", res)
		}
	})
}
