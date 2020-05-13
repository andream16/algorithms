package single

import "testing"

func TestSingleNonDuplicate(t *testing.T) {
	t.Run("it should return 2", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
		if 2 != nonDuplicate {
			t.Fatalf("expected 2, got %d", nonDuplicate)
		}
	})
	t.Run("it should return 10", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11})
		if 10 != nonDuplicate {
			t.Fatalf("expected 10, got %d", nonDuplicate)
		}
	})
	t.Run("it should return 1", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{1, 3, 3, 7, 7, 10, 10, 11, 11})
		if 1 != nonDuplicate {
			t.Fatalf("expected 1, got %d", nonDuplicate)
		}
	})
	t.Run("it should return 13", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{1, 1, 3, 3, 7, 7, 10, 10, 11, 11, 13})
		if 13 != nonDuplicate {
			t.Fatalf("expected 13, got %d", nonDuplicate)
		}
	})
	t.Run("it should return 1", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{1})
		if 1 != nonDuplicate {
			t.Fatalf("expected 1, got %d", nonDuplicate)
		}
	})
	t.Run("it should return 3", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{1, 1, 3, 4, 4})
		if 3 != nonDuplicate {
			t.Fatalf("expected 3, got %d", nonDuplicate)
		}
	})
	t.Run("it should return 0 because there are no duplicates", func(t *testing.T) {
		nonDuplicate := singleNonDuplicate([]int{1, 1, 3, 3, 4, 4})
		if 0 != nonDuplicate {
			t.Fatalf("expected 0, got %d", nonDuplicate)
		}
	})
}
