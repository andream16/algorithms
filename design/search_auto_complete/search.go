package main

import (
	"fmt"
	"sort"
)

type AutocompleteSystem struct {
	trie    *TrieNode
	history []rune
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	return AutocompleteSystem{
		trie:    buildTrie(sentences, times),
		history: []rune{},
	}
}

// TODO finish for next searches.
func (this *AutocompleteSystem) Input(c byte) []string {
	this.history = append(this.history, rune(c))
	return this.trie.lookup(this.history)
}

type TrieNode struct {
	isRoot     bool
	isTerminal bool
	val        rune
	strVal     string
	next       map[rune]*TrieNode
	hotness    int
}

type HotSentence struct {
	val     string
	hotness int
}

type ByHotness []HotSentence

func (b ByHotness) Len() int      { return len(b) }
func (b ByHotness) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b ByHotness) Less(i, j int) bool {
	if b[i].hotness == b[j].hotness {
		for k := 0; k < len(b[k].val); k++ {
			if b[i].val[k] == b[j].val[k] {
				continue
			}
			return b[i].val[k] < b[j].val[k]
		}
	}
	return b[i].hotness > b[j].hotness
}

func (t *TrieNode) lookup(pre []rune) []string {
	var (
		hotSentences = []HotSentence{}
		sentences    = []string{}
		curr         = *t
		currWord     string
	)
	for _, c := range pre {
		next, ok := curr.next[c]
		if !ok {
			return nil
		}
		currWord += string(c)
		curr = *next
	}

	if curr.isTerminal {
		return []string{currWord}
	} else {
		for _, w := range curr.next {
			hotSentences = append(hotSentences, getWords(w, currWord)...)
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

func getWords(node *TrieNode, prefix string) []HotSentence {
	prefix += string(node.val)
	if node.isTerminal {
		return []HotSentence{{val: prefix, hotness: node.hotness}}
	}
	hotSentences := []HotSentence{}
	for _, next := range node.next {
		hotSentences = append(hotSentences, getWords(next, prefix)...)
	}
	return hotSentences
}

func buildTrie(sentences []string, hotness []int) *TrieNode {
	trieNode := TrieNode{
		isRoot: true,
		next:   map[rune]*TrieNode{},
	}

	for i, s := range sentences {
		nextNode := &trieNode
		for j, c := range s {
			node, ok := nextNode.next[c]
			if ok {
				nextNode = node
				continue
			} else {
				node = &TrieNode{}
			}

			node.val = c
			node.next = map[rune]*TrieNode{}
			node.strVal = string(c)

			if j == len(s)-1 {
				node.isTerminal = true
				node.hotness = hotness[i]
			}

			nextNode.next[c] = node
			nextNode = node
		}
	}

	return &trieNode
}

func main() {
	autoComplete := Constructor([]string{"i love you", "island", "ironman", "i love leetcode"}, []int{5, 3, 2, 2})
	fmt.Println(autoComplete.Input('i')) // ["i love you", "island","i love leetcode"]
	fmt.Println(autoComplete.Input(' ')) // ["i love you", "i love leetcode"]
	fmt.Println(autoComplete.Input('a')) // []
	fmt.Println(autoComplete.Input('#')) // []

	//trie := buildTrie([]string{"i love you", "island", "ironman", "i love leetcode"}, []int{5, 3, 2, 2})
	//fmt.Println(trie.lookup([]rune{'i'}))
}
