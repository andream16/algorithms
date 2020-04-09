package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func main() {}
