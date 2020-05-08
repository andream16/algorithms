package main

import (
	"fmt"
	"sort"
)

type AutocompleteSystem struct {
	trie    *TrieNode
	history string
}

type TrieNode struct {
	isWord   bool
	val      rune
	children map[rune]*TrieNode
	hotness  int
}

type HotSentence struct {
	val     string
	hotness int
}

type ByHotness []HotSentence

type QueueNode struct {
	node   *TrieNode
	prefix string
}

type Queue struct {
	nodes []*QueueNode
}

// T: O(sentence*letters)
// S: (sentences*letters*children)
func Constructor(sentences []string, times []int) AutocompleteSystem {
	trieNode := TrieNode{
		children: map[rune]*TrieNode{},
	}
	for i, s := range sentences {
		trieNode.upsert(s, times[i])
	}
	return AutocompleteSystem{
		trie: &trieNode,
	}
}

// T: O(prefix + (children*letters) + matches(log matches))
func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.trie.upsert(this.history, 1)
		this.history = ""
		return make([]string, 0)
	}
	this.history += string(c)
	return this.trie.lookup(this.history)
}

func (b ByHotness) Len() int      { return len(b) }
func (b ByHotness) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByHotness) Less(i, j int) bool {
	if b[i].hotness == b[j].hotness {
		return b[i].val < b[j].val
	}
	return b[i].hotness > b[j].hotness
}

func (t *TrieNode) lookup(prefix string) []string {

	var (
		curr     = *t
		currWord string
	)

	for _, c := range prefix {
		next, ok := curr.children[c]
		if !ok {
			return make([]string, 0)
		}
		currWord += string(c)
		curr = *next
	}

	var (
		hotSentences []HotSentence
		sentences    []string
	)

	q := &Queue{}
	q.enqueue(&QueueNode{
		node: &curr,
	})

	for !q.empty() {
		n := q.dequeue()
		if n.node.isWord {
			hotSentences = append(hotSentences, HotSentence{
				val:     prefix + n.prefix,
				hotness: n.node.hotness,
			})
		}
		for _, child := range n.node.children {
			q.enqueue(&QueueNode{
				node:   child,
				prefix: n.prefix + string(child.val),
			})
		}
	}

	sort.Sort(ByHotness(hotSentences))

	for i := 0; i < len(hotSentences); i++ {
		if i == 3 {
			break
		}
		sentences = append(sentences, hotSentences[i].val)
	}

	return sentences
}

func (t *TrieNode) upsert(sentence string, hotness int) {

	curr := t

	for i, c := range sentence {
		var nextNode *TrieNode
		node, ok := curr.children[c]
		if ok {
			nextNode = node
		} else {
			nextNode = &TrieNode{
				val:      c,
				children: map[rune]*TrieNode{},
			}
		}

		if i == len(sentence)-1 {
			nextNode.isWord = true
			if nextNode.hotness == 0 {
				nextNode.hotness = hotness
			} else {
				nextNode.hotness++
			}
		}

		curr.children[c] = nextNode
		curr = nextNode
	}

}

func (q *Queue) enqueue(node *QueueNode) {
	q.nodes = append(q.nodes, node)
}

func (q *Queue) dequeue() *QueueNode {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node
}

func (q *Queue) empty() bool {
	return len(q.nodes) == 0
}

func main() {

	// ["AutocompleteSystem","input","input","input","input","input","input","input","input","input","input","input","input"]
	//[[["i love you","island","iroman","i love leetcode"],[5,3,2,2]],["i"],[" "],["a"],["#"],["i"],[" "],["a"],["#"],["i"],[" "],["a"],["#"]]

	autoComplete := Constructor([]string{"abc", "abbc", "a"}, []int{3, 3, 3})
	fmt.Println(autoComplete.Input('b')) // []
	fmt.Println(autoComplete.Input('c')) // []
	fmt.Println(autoComplete.Input('#')) // []
	fmt.Println(autoComplete.Input('b')) // ["bc"]
	fmt.Println(autoComplete.Input('c')) // ["bc"]
	fmt.Println(autoComplete.Input('#')) // []
	fmt.Println(autoComplete.Input('a')) // ["a", "abbc", "abc"]
	fmt.Println(autoComplete.Input('b')) // ["abbc", "abc"]
	fmt.Println(autoComplete.Input('c')) // ["abc"]
	fmt.Println(autoComplete.Input('#')) // []
	fmt.Println(autoComplete.Input('a')) // ["abc", "a", "abbc"]
	fmt.Println(autoComplete.Input('b')) // ["abc", "abbc"]
	fmt.Println(autoComplete.Input('c')) // ["abc"]
	fmt.Println(autoComplete.Input('#')) // []

}
