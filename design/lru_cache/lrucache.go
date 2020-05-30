package lrucache

type ListNode struct {
	Key  int
	Val  int
	Next *ListNode
	Prev *ListNode
}

type LRUCache struct {
	Capacity  int
	KeysNodes map[int]*ListNode
	Head      *ListNode
	Tail      *ListNode
}

// We can solve this problem using an map to hold key and values as usual and a
// linked list with dummy head and tail to update the order in constant time.
// It's important to keep in mind that when we reach the capacity, we need to make space for the new element.
// The core of the logic for Get & Put are add and remove.
//
// T: O(1)
// S: O(capacity)
func Constructor(capacity int) LRUCache {
	head, tail := &ListNode{}, &ListNode{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity:  capacity,
		KeysNodes: make(map[int]*ListNode, capacity),
		Head:      head,
		Tail:      tail,
	}
}

// We look for the key in the dictionary. If it doesn't exists we are done.
// If it exists, we need to remove it and add it again.
//
// T: O(1)
// S: O(1)
func (this *LRUCache) Get(key int) int {
	v, ok := this.KeysNodes[key]
	if !ok {
		return -1
	}
	this.remove(v)
	this.add(v)
	return v.Val
}

// We look for the key in the dictionary. If it exists, it means that we need to update the respective node with the new
// value and making it the most recent element. We do so by removing it and re-adding it.
// If it doesn't exist, we need to check if we reached the capacity. If we did reach it, we need to delete the least
// used element from both list and map before adding a new one.
// Otherwise, we simply add a new one to both list and map.
//
// T: O(1)
// S: O(1)
func (this *LRUCache) Put(key int, value int) {
	v, ok := this.KeysNodes[key]
	if ok {
		this.remove(v)
		v.Val = value
		this.add(v)
		return
	}

	if this.isFull() {
		delete(this.KeysNodes, this.Tail.Prev.Key)
		this.remove(this.Tail.Prev)
	}

	newNode := &ListNode{Key: key, Val: value}
	this.KeysNodes[key] = newNode
	this.add(newNode)
}

// to add a node in constant time, we need to rely on the head.
// The least used element is always head's next.
// So, we update like following to add n1:
// - h -> n0 -> t
// - h -> n1 -> n0 -> t
// This works also if we don't have n0.
func (this *LRUCache) add(node *ListNode) {
	headNext := this.Head.Next
	this.Head.Next = node
	node.Next = headNext
	node.Prev = this.Head
	headNext.Prev = node
}

//to remove a node in constant time, we need to update its prev and next to rely on each other.
func (this *LRUCache) remove(node *ListNode) {
	prev, next := node.Prev, node.Next
	prev.Next = next
	next.Prev = prev
}

// checks if we reached the cache max size.
func (this *LRUCache) isFull() bool {
	return this.Capacity == len(this.KeysNodes)
}
