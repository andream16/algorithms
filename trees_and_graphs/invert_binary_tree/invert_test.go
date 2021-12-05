package invert_binary_tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInvert(t *testing.T) {
	for i, tt := range []struct {
		givenBinaryTree    *TreeNode
		expectedBinaryTree *TreeNode
	}{
		{
			givenBinaryTree: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 9,
					},
				},
			},
			expectedBinaryTree: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 7,
					Left: &TreeNode{
						Val: 9,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 1,
					},
				},
			},
		},
		{
			givenBinaryTree: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			expectedBinaryTree: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		{
			givenBinaryTree:    &TreeNode{},
			expectedBinaryTree: &TreeNode{},
		},
	} {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			root := invertTree(tt.givenBinaryTree)
			if !reflect.DeepEqual(tt.expectedBinaryTree, root) {
				t.Fatalf("expected tree %v, got %v", tt.expectedBinaryTree, root)
			}
		})
	}
}
