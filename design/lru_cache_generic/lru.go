package lru_cache_generic

type (
	LRUCache[K, V comparable] struct {
		capacity uint
		head     *Node[K, V]
		tail     *Node[K, V]
		cache    map[K]*Node[K, V]
	}
	Node[K, V comparable] struct {
		value V
		key   K
		next  *Node[K, V]
		prev  *Node[K, V]
	}
)

func New[K, V comparable](capacity uint) *LRUCache[K, V] {
	tail, head := &Node[K, V]{}, &Node[K, V]{}
	head.next = tail
	tail.prev = head

	return &LRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*Node[K, V], capacity),
		head:     head,
		tail:     tail,
	}
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	n, ok := c.cache[key]
	if ok {
		c.remove(n)
		n.value = value
		c.add(n)
		return
	}

	if len(c.cache) >= int(c.capacity) {
		delete(c.cache, c.tail.prev.key)
		c.remove(c.tail.prev)
	}

	newNode := &Node[K, V]{key: key, value: value}
	c.cache[key] = newNode
	c.add(newNode)
}

func (c *LRUCache[K, V]) Get(key K) (v V, ok bool) {
	n, ok := c.cache[key]
	if !ok {
		return v, false
	}
	c.remove(n)
	c.add(n)
	return n.value, true
}

func (c *LRUCache[K, V]) remove(n *Node[K, V]) {
	prev, next := n.prev, n.next
	prev.next = next
	next.prev = prev
}

func (c *LRUCache[K, V]) add(n *Node[K, V]) {
	headNext := c.head.next
	c.head.next = n
	n.next = headNext
	n.prev = c.head
	headNext.prev = n
}
