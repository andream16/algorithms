package oddeven

import "testing"

func TestOddEven(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		input := &ListNode{
			Val:  1,
			Next: &ListNode{
				Val:  2,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  4,
						Next: &ListNode{
							Val:  5,
						},
					},
				},
			},
		}
		output := oddEvenList(input)
		if 1 != output.Val {
			t.Fatalf("expected 1, got %d", output.Val)
		}
		output = output.Next
		if 3 != output.Val {
			t.Fatalf("expected 3, got %d", output.Val)
		}
		output = output.Next
		if 5 != output.Val {
			t.Fatalf("expected 5, got %d", output.Val)
		}
		output = output.Next
		if 2 != output.Val {
			t.Fatalf("expected 2, got %d", output.Val)
		}
		output = output.Next
		if 4 != output.Val {
			t.Fatalf("expected 4, got %d", output.Val)
		}
		if output.Next != nil {
			t.Fatal("expected next nil")
		}
	})
	t.Run("test case 2", func(t *testing.T) {
		input := &ListNode{
			Val:  2,
			Next: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{
							Val:  6,
							Next: &ListNode{
								Val:  4,
								Next: &ListNode{
									Val:  7,
								},
							},
						},
					},
				},
			},
		}
		output := oddEvenList(input)
		if 2 != output.Val {
			t.Fatalf("expected 2, got %d", output.Val)
		}
		output = output.Next
		if 3 != output.Val {
			t.Fatalf("expected 3, got %d", output.Val)
		}
		output = output.Next
		if 6 != output.Val {
			t.Fatalf("expected 6, got %d", output.Val)
		}
		output = output.Next
		if 7 != output.Val {
			t.Fatalf("expected 7, got %d", output.Val)
		}
		output = output.Next
		if 1 != output.Val {
			t.Fatalf("expected 1, got %d", output.Val)
		}
		output = output.Next
		if 5 != output.Val {
			t.Fatalf("expected 5, got %d", output.Val)
		}
		output = output.Next
		if 4 != output.Val {
			t.Fatalf("expected 4, got %d", output.Val)
		}
		if output.Next != nil {
			t.Fatal("expected next nil")
		}
	})
}
