package reverse

import "testing"

func TestReverseList(t *testing.T) {
	input := &ListNode{
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
	output := reverseList(input)
	if 5 != output.Val {
		t.Fatalf("expected 5, got %d", output.Val)
	}
	output = output.Next
	if 4 != output.Val {
		t.Fatalf("expected 4, got %d", output.Val)
	}
	output = output.Next
	if 3 != output.Val {
		t.Fatalf("expected 3, got %d", output.Val)
	}
	output = output.Next
	if 2 != output.Val {
		t.Fatalf("expected 2, got %d", output.Val)
	}
	output = output.Next
	if 1 != output.Val {
		t.Fatalf("expected 1, got %d", output.Val)
	}
	if nil != output.Next {
		t.Fatalf("expected next to be nil but has value %d", output.Next.Val)
	}
}

func TestReverseListNonOpt(t *testing.T) {
	input := &ListNode{
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
	output := reverseListNonOpt(input)
	if 5 != output.Val {
		t.Fatalf("expected 5, got %d", output.Val)
	}
	output = output.Next
	if 4 != output.Val {
		t.Fatalf("expected 4, got %d", output.Val)
	}
	output = output.Next
	if 3 != output.Val {
		t.Fatalf("expected 3, got %d", output.Val)
	}
	output = output.Next
	if 2 != output.Val {
		t.Fatalf("expected 2, got %d", output.Val)
	}
	output = output.Next
	if 1 != output.Val {
		t.Fatalf("expected 1, got %d", output.Val)
	}
	if nil != output.Next {
		t.Fatalf("expected next to be nil but has value %d", output.Next.Val)
	}
}
