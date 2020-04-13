package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// T: O(max(|l1|,|l2|))
// S: O(max(|l1|,|l2|)) + 1
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		res   = ListNode{}
		curr  = &res
		carry int
	)

	if l1 == nil && l2 == nil {
		return &res
	}

	for {

		var v1, v2 int

		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}

		val := v1 + v2 + carry
		carry = val / 10
		curr.Val = val % 10

		if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) {
			if carry > 0 {
				n := &ListNode{Val: carry}
				curr.Next = n
				curr = curr.Next
			}
			break
		}

		curr.Next = &ListNode{}
		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return &res
}
