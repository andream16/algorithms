package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// To merge two sorted linked lists, we need to pay attention to the order of their values.
// We iterate both lists at the same time, but go to next only on the list that has a current node value smaller
// than the one in the other current list. We have to handle the cases when one or the other list are empty.
//
// T: O(max(|l1|,|l2|))
// S: |l1|+|l2|
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		res  = ListNode{}
		curr = &res
	)

	for l1 != nil || l2 != nil {

		n := ListNode{}

		switch {
		case l1 != nil && l2 != nil:
			if l1.Val < l2.Val {
				n.Val = l1.Val
				l1 = l1.Next
				break
			}
			n.Val = l2.Val
			l2 = l2.Next
		case l1 != nil && l2 == nil:
			n.Val = l1.Val
			l1 = l1.Next
		case l1 == nil && l2 != nil:
			n.Val = l2.Val
			l2 = l2.Next
		}

		curr.Next = &n
		curr = curr.Next
	}

	return res.Next
}
