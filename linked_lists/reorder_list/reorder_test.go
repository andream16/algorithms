package reorderlist

import "testing"

func TestReorderList(t *testing.T) {
	t.Run("should successfully reorder a list", func(t *testing.T) {
		l := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		}
		reorderList(l)
		expected := &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
			},
		}

		for l != nil && expected != nil {
			if expected.Val != l.Val {
				t.Fatalf("expected %d, got %d instead", expected.Val, l.Val)
			}
			l = l.Next
			expected = expected.Next
		}
	})
}
