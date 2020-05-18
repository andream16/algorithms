package average

import "testing"

func TestAverage(t *testing.T) {
	m := Constructor(3)
	if n := m.Next(1); n != float64(1) {
		t.Fatalf("expected 1, got %f", n)
	}
	if n := m.Next(10); n != 5.5 {
		t.Fatalf("expected 5.5, got %f", n)
	}
	if n := m.Next(3); n != 4.666667 {
		t.Fatalf("expected 4.666667, got %f", n)
	}
	if n := m.Next(5); n != 6 {
		t.Fatalf("expected 6, got %f", n)
	}
}
