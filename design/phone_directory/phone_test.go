package phonedirectory

import "testing"

func TestPhoneDirectory(t *testing.T) {
	t.Run("testcase 1", func(t *testing.T) {
		dir := Constructor(3)
		if n := dir.Get(); 0 != n {
			t.Fatalf("expected 0, got %d", n)
		}
		if n := dir.Get(); 1 != n {
			t.Fatalf("expected 1, got %d", n)
		}
		if !dir.Check(2) {
			t.Fatal("expected 2 to be available")
		}
		if n := dir.Get(); 2 != n {
			t.Fatalf("expected 2, got %d", n)
		}
		if dir.Check(2) {
			t.Fatal("expected 2 to not be available")
		}
		dir.Release(2)
		if !dir.Check(2) {
			t.Fatal("expected 2 to be available")
		}
	})
	t.Run("testcase 2", func(t *testing.T) {
		dir := Constructor(2)
		dir.Release(1)
		if n := dir.Get(); 0 != n {
			t.Fatalf("expected 0, got %d", n)
		}
		if !dir.Check(1) {
			t.Fatal("expected 1 to be available")
		}
		if !dir.Check(1) {
			t.Fatal("expected 1 to be available")
		}
		dir.Release(1)
		if !dir.Check(1) {
			t.Fatal("expected 1 to be available")
		}
		if n := dir.Get(); 1 != n {
			t.Fatalf("expected 1, got %d", n)
		}
		if dir.Check(0) {
			t.Fatal("expected 0 to not be available")
		}
		if dir.Check(1) {
			t.Fatal("expected 1 to not be available")
		}
		if dir.Check(1) {
			t.Fatal("expected 1 to not be available")
		}
	})
}
