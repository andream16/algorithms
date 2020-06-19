package reorderlist

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack struct {
	Nodes []*ListNode
}

func (q *Stack) push(node *ListNode) {
	q.Nodes = append([]*ListNode{node}, q.Nodes...)
}

func (q *Stack) pop(ctr int) *ListNode {
	var n *ListNode
	if ctr%2 == 0 {
		n = q.Nodes[len(q.Nodes)-1]
		q.Nodes = q.Nodes[:len(q.Nodes)-1]
	} else {
		n = q.Nodes[0]
		q.Nodes = q.Nodes[1:]
	}
	return n
}

func (q *Stack) empty() bool {
	return len(q.Nodes) == 0
}

// TODO REVIEW
// T: O(|list|) + O(|list|)
// S: O(|list|)
func reorderList(head *ListNode) {

	var (
		n   = head
		q   = &Stack{}
		ctr int
	)

	for n != nil {
		q.push(n)
		n = n.Next
	}

	for !q.empty() {
		if len(q.Nodes) == 1 {
			head.Next = nil
		} else {
			head.Next = &ListNode{}
		}
		head.Val = q.pop(ctr).Val
		head = head.Next
		ctr++
	}

}
