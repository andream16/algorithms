package duplicate

import "testing"

func TestFindDuplicate(t *testing.T) {
	if res := findDuplicate([]int{1, 2, 3, 4, 1, 5, 6, 7}); 1 != res {
		t.Fatalf("expected 1, got %d instead", res)
	}
}

func TestFindDuplicateNaive(t *testing.T) {
	if res := findDuplicateNaive([]int{1, 2, 3, 4, 1, 5, 6, 7}); 1 != res {
		t.Fatalf("expected 1, got %d instead", res)
	}
}

func BenchmarkFindDuplicate10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDuplicate([]int{1, 2, 3, 4, 1, 5, 6, 7})
	}
}

// BenchmarkFindDuplicate10000
// BenchmarkFindDuplicate10000-8           210498027                5.69 ns/op

// BenchmarkFindDuplicateNaive10000
// BenchmarkFindDuplicateNaive10000-8      15626240                72.9 ns/op
func BenchmarkFindDuplicateNaive10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDuplicateNaive([]int{1, 2, 3, 4, 1, 5, 6, 7})
	}
}
