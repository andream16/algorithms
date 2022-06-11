package multiply_strings

import "testing"

func TestMultiply(t *testing.T) {
	for _, tt := range []struct {
		n1       string
		n2       string
		expected string
	}{
		{n1: "0", n2: "1", expected: "0"},
		{n1: "1", n2: "0", expected: "0"},
		{n1: "1", n2: "10", expected: "10"},
		{n1: "12", n2: "4567", expected: "54804"},
		{n1: "121212 ", n2: "11010", expected: "1334544120"},
	} {
		if res := multiply(tt.n1, tt.n2); res != tt.expected {
			t.Fatalf("expected multiplication between %s and %s to return %s. Got %s instead.", tt.n1, tt.n2, tt.expected, res)
		}
	}
}
