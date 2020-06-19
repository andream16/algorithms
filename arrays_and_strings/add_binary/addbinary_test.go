package addbinary

import "testing"

func TestAddBinary(t *testing.T) {
	for _, tc := range []struct {
		a        string
		b        string
		expected string
	}{
		{a: "11", b: "1", expected: "100"},
		{a: "1010", b: "1011", expected: "10101"},
	} {
		if res := addBinary(tc.a, tc.b); tc.expected != res {
			t.Fatalf("expected %s, got %s instead", tc.expected, res)
		}
	}
}
