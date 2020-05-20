package reverse

type ListNode struct {
	Val  int
	Next *ListNode
}

// Simply swap them as we iterate.
// The first time, prev has to be nil otherwise we'll have an extra ListNode at the end.
//
// T: O(n)
// S: O(1)
func reverseList(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	prev, curr := &ListNode{}, head
	prev = nil

	for nil != curr {
		tmp := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmp
	}

	return prev
}

// Copy all the values as we iterate and append them the other way around.
//
// T: O(n)
// S: O(1)
func reverseListNonOpt(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	var queue []int

	for nil != head {
		queue = append(queue, head.Val)
		head = head.Next
	}

	res := ListNode{}
	curr := &res

	for i := len(queue) - 1; i >= 0; i-- {
		curr.Val = queue[i]
		if i > 0 {
			curr.Next = &ListNode{}
			curr = curr.Next
		}
	}

	return &res
}
