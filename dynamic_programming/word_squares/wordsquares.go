package main

import (
	"fmt"
	"strings"
)

type TrieNode struct {
	Val   string
	Nodes map[string]*TrieNode
}

func buildTrie(words []string) *TrieNode {

	node := TrieNode{
		Val: "*",
	}
	currNode := &node

	for _, word := range words {
		currNode = &node
		for _, char := range strings.Split(word, "") {
			n, ok := currNode.Nodes[char]
			if ok {
				currNode = n
				continue
			}
			currNode.Val = char
			nextNode := &TrieNode{}
			if currNode.Nodes == nil {
				currNode.Nodes = map[string]*TrieNode{}
			}
			currNode.Nodes[char] = nextNode
			currNode = nextNode
		}
		node.Nodes[string(word[0])] = currNode
	}
	return &node
}

func wordSquares(words []string) [][]string {
	var (
		results = [][]string{}
	)

	_ = buildTrie(words)

	// buildTrie

	// getResults
	//    getPrefixAtCurrentIndex
	//    if ok, continue recursively

	return results
}

func main() {
	fmt.Println(wordSquares([]string{"area", "lead", "wall", "lady", "ball"}))
}
