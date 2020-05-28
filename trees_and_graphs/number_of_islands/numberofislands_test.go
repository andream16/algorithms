package island

import "testing"

func TestNumIslands(t *testing.T) {
	t.Run("it should return 1 island", func(t *testing.T) {
		if n := numIslands([][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}); 1 != n {
			t.Fatalf("expected 1 island but found %d", n)
		}
	})
	t.Run("it should return 3 islands", func(t *testing.T) {
		if n := numIslands([][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}); 3 != n {
			t.Fatalf("expected 3 islands but found %d", n)
		}
	})
	t.Run("it should return 1 island", func(t *testing.T) {
		if n := numIslands([][]byte{
			{'1', '1', '1'},
			{'0', '1', '0'},
			{'1', '1', '1'},
		}); 1 != n {
			t.Fatalf("expected 1 island but found %d", n)
		}
	})
}
