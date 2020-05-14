package trie

type Trie struct {
	Val      rune
	Final    bool
	Children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{
		Children: map[rune]*Trie{},
	}
}

// T: O(n)
// S: O(n*n-1)
/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	node := this
	for _, c := range word {
		child, ok := node.Children[c]
		if ok {
			node = child
			continue
		}
		newNode := &Trie{
			Val:      c,
			Children: map[rune]*Trie{},
		}
		node.Children[c] = newNode
		node = newNode
	}
	node.Final = true
}

// T: O(n)
// S: O(1)
/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, c := range word {
		child, ok := node.Children[c]
		if !ok {
			return false
		}
		node = child
	}
	return node.Final
}

// T: O(n)
// S: O(1)
/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, c := range prefix {
		child, ok := node.Children[c]
		if !ok {
			return false
		}
		node = child
	}
	return true
}
