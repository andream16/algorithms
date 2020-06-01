package deepcopy

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// Cloning different nodes values is straightforward. The problem comes in copying correctly the random references
// as their nodes may not have been visited yet. To solve this, we adopt a greedy approach.
// While we copy the nodes, leaving the random reference outside, we map each old node to its respective new node.
// When we are done, we fill the missing random information by looking into the map for the new node that maps the
// current random.
//
// T: O(n)
// S: O(n)
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var (
		oldToNew = map[*Node]*Node{}
		deepCopy = Node{}
		curr     = &deepCopy
	)

	for nil != head {
		curr.Val = head.Val
		if nil != head.Next {
			curr.Next = &Node{}
		}
		oldToNew[head] = curr
		head = head.Next
		curr = curr.Next
	}

	for oldNode, newNode := range oldToNew {
		if nil == oldNode.Random {
			continue
		}
		rnd, ok := oldToNew[oldNode.Random]
		if !ok {
			continue
		}
		newNode.Random = rnd
	}

	return &deepCopy
}
