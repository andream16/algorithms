package insertdelete

import "testing"

func TestInsertDeleteGetRandom(t *testing.T) {
	t.Run("scenario 1", func(t *testing.T) {
		set := Constructor()
		if true != set.Insert(1) {
			t.Fatalf("expected 'true' but got 'false' when inserting '1'")
		}
		if false != set.Remove(2) {
			t.Fatalf("expected 'false' but got 'true' when removing '2'")
		}
		if true != set.Insert(2) {
			t.Fatalf("expected 'true' but got 'false' when inserting '2'")
		}
		rand := set.GetRandom()
		if _, ok := map[int]struct{}{
			1: {},
			2: {},
		}[rand]; !ok {
			t.Fatalf("expected random number to be '[1, 2]', got %d instead", rand)
		}
		if true != set.Remove(1) {
			t.Fatalf("expected 'true' but got 'false' when removing '1'")
		}
		if false != set.Insert(2) {
			t.Fatalf("expected 'false' but got 'true' when inserting '2'")
		}
		if _, ok := map[int]struct{}{
			2: {},
		}[rand]; !ok {
			t.Fatalf("expected random number to be '[2]', got %d instead", rand)
		}
	})
	t.Run("scenario 2", func(t *testing.T) {
		set := Constructor()
		if false != set.Remove(0) {
			t.Fatalf("expected 'false' but got 'true' when removing '0'")
		}
		if false != set.Remove(0) {
			t.Fatalf("expected 'false' but got 'true' when removing '0'")
		}
		if true != set.Insert(0) {
			t.Fatalf("expected 'true' but got 'false' when inserting '0'")
		}
		rand := set.GetRandom()
		if _, ok := map[int]struct{}{
			0: {},
		}[rand]; !ok {
			t.Fatalf("expected random number to be '[0]', got %d instead", rand)
		}
		if true != set.Remove(0) {
			t.Fatalf("expected 'true' but got 'false' when removing '0'")
		}
		if true != set.Insert(0) {
			t.Fatalf("expected 'true' but got 'false' when inserting '0'")
		}
	})
}
