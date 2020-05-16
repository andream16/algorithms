package oddeven

type ListNode struct {
	Val int
	Next *ListNode
}

// TODO review this
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddEnd := head

	evenHead := head.Next
	evenEnd := evenHead

	for evenEnd != nil && evenEnd.Next != nil {
		oddEnd.Next = evenEnd.Next
		oddEnd = oddEnd.Next

		evenEnd.Next = oddEnd.Next
		evenEnd = evenEnd.Next
	}

	oddEnd.Next = evenHead

	return head
}