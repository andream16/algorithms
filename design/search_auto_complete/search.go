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

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.trie.addOrUpdateSentence(this.history)
		this.history = []rune{}
		return make([]string, 0)
	}
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
		for k := 0; k < len(min(b[i].val, b[j].val)); k++ {
			if b[i].val[k] == b[j].val[k] {
				continue
			}
			return b[i].val[k] < b[j].val[k]
		}
	}
	return b[i].hotness > b[j].hotness
}

func min(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s2
	}
	return s1
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
		hotSentences = append(hotSentences, HotSentence{
			val:     currWord,
			hotness: curr.hotness,
		})
		if len(curr.next) == 0 {
			return []string{currWord}
		}
	}

	for _, w := range curr.next {
		hotSentences = append(hotSentences, getWords(w, currWord)...)
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

func (t *TrieNode) addOrUpdateSentence(sentence []rune) {

	nextNode := *t

	for i, c := range sentence {
		node, ok := nextNode.next[c]
		if ok {
			if node.isTerminal && len(node.next) == 0 {
				node.hotness++
			}
			nextNode = *node
			continue
		} else {
			node = &TrieNode{}
		}
		node.val = c
		node.next = map[rune]*TrieNode{}
		node.strVal = string(c)

		if i == len(sentence)-1 {
			node.isTerminal = true
			node.hotness = 1
		}

		nextNode.next[c] = node
		nextNode = *node
	}

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
				if j == len(s)-1 {
					node.isTerminal = true
					node.hotness = hotness[i]
				}
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
