package maxpathsum

import "testing"

func TestMaxPathSum(t *testing.T) {
	t.Run("should return the correct max path sum", func(t *testing.T) {
		node := &TreeNode{
			Val:   -10,
			Left:  &TreeNode{
				Val:   9,
			},
			Right: &TreeNode{
				Val:   20,
				Left:  &TreeNode{
					Val:   15,
				},
				Right: &TreeNode{
					Val:   7,
				},
			},
		}
		if sum := maxPathSum(node); 42 != sum {
			t.Fatalf("expected max path sum to be 42, got %d instead", sum)
		}
	})
}
