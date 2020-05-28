package diameter

import "testing"

func TestDiameterOfABinaryTree(t *testing.T) {
	tree := &TreeNode{
		Left: &TreeNode{
			Left:  &TreeNode{},
			Right: &TreeNode{},
		},
		Right: &TreeNode{},
	}
	if diameter := diameterOfBinaryTree(tree); 3 != diameter {
		t.Fatalf("expected diameter 3, got %d", diameter)
	}
}
